package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	businessTypes "github.com/quinlanmorake/verisart-go/types"
	
	database "github.com/quinlanmorake/verisart-go/database"
	userPackage "github.com/quinlanmorake/verisart-go/user"		
)

func CreateTransfer(user businessTypes.User, certificateId coreTypes.String, receiver coreTypes.EmailAddress) (businessTypes.Transfer, coreTypes.Result) {
	var certificate businessTypes.Certificate

	// First validate that we are able to create a transfer	
	if _certificate, validateTransferableResult := validateBeforeCreatingTransfer(user, certificateId, receiver); validateTransferableResult.IsNotOk() {
		return businessTypes.NewEmptyTransfer(), validateTransferableResult
	} else {
		// We need to update data on the certificate object as well
		certificate = _certificate
	}

	/*
   This is a tricky bit, but we will handle in a simple fashion. 
   First we update the certificate, and add the transfer to it. In that manner, we in practicality, avoid the race condition of 2 people simultaneously attempting to create a transfer
   The limits of this vary, but with an in memory database that has a locking mechanism, we know we are pretty close to ensuring this.

   We then create the transfer; and if something fails, we undo the edit on the the certificate record, and restore the its transfer status to "None"

   There are 2 other possible race conditions. 
    Once the transfer is created; the receiver clicks accept, and the sender clicks edit / delete; Here again, we are relying on the in-memory database.
    If accept happens before edit / delete; edit / delete will fail, because the sender is no longer the owner of the certificate
    If edit before accept, edit will succeed (this is by design)
    If delete happens before accept; accept will fail, noting that the certificate no longer exists (this is by design)
  */
	transfer := businessTypes.NewTransfer(user.Id, certificate.Id, receiver)
	
	certificate.Transfer = businessTypes.TransferState {
		CreatedAt: transfer.CreatedAt,
		CreatedBy: user.Id,
		Status: transfer.Status,
	}
	
	if updateCertificateResult := database.Update(&certificate); updateCertificateResult.IsNotOk() {
		return businessTypes.NewEmptyTransfer(), updateCertificateResult
	}

	if addTransferResult := database.Add(&transfer); addTransferResult.Result.IsNotOk() {
		// Rollback certificate
		// ALERT: possible data inconsistency here, so on an error adding a transfer, application support needs to check the certificate record.
		// Should many such errors occur, then a solution around such should be sought; simple one would be to attaching the transfer as a property on the certificate record if NoSql is being used
		// wrap in a transaction in a stored procedure if SQL is being used
		// With in memory db, this shouldn't be a problem unless the app runs out of memory or other such corruption, at which poitn data inconsistency is inevitable.
		certificate.Transfer = businessTypes.TransferState{}
		database.Update(&certificate)

		return businessTypes.NewEmptyTransfer(), addTransferResult.Result
	} else {
		return transfer,  coreTypes.NewSuccessResult()
	}
}

func validateBeforeCreatingTransfer(user businessTypes.User, certificateId coreTypes.String, receiver coreTypes.EmailAddress) (certificate businessTypes.Certificate, result coreTypes.Result) {
	// Does sender exist
	if senderExists, checkSenderExistsResult := userPackage.UserWithEmailExists(receiver); checkSenderExistsResult.IsNotOk()  {
		result = checkSenderExistsResult
		return
	} else {
		if senderExists == false {
			result = coreTypes.NewResultFromErrorCode(errorCodes.NO_USER_WITH_THE_EMAIL_EXISTS)
			return
		}
	}
	
	// Does certificate exist
	if _certificate, loadCertificateResult := GetCertificateById(certificateId); loadCertificateResult.IsNotOk() {
		result = loadCertificateResult
		return
	} else {
		certificate = _certificate
	}

	// Is the user the owner of the certificate
	if certificate.OwnerId.IsEqual(user.Id) == false {
		result = coreTypes.NewResultFromErrorCode(errorCodes.INSUFFICIENT_PERMISSIONS_TO_MODIFY_THE_REQUESTED_CERTIFICATE)
		return
	}

	// Is the user trying to transfer the certificate to themself
	if user.Email.IsEqual(receiver) {
		result = coreTypes.NewResultFromErrorCode(errorCodes.UNABLE_TO_TRANSFER_TO_YOURSELF)
		return
	}
	
	// Does the certificate have an associated transfer
	if pendingTransferAlreadyExists, checkTransferResult := doesPendingTransferForCertificateExist(certificate); checkTransferResult.IsNotOk() {
		result = checkTransferResult
		return		
	} else {
		if pendingTransferAlreadyExists {
			result = coreTypes.NewResultFromErrorCode(errorCodes.CERTIFICATE_ALREADY_HAS_PENDING_TRANSFER)
			return
		}
	}

	result = coreTypes.NewSuccessResult()
	return
}