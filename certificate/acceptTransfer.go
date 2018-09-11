package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	businessTypes "github.com/quinlanmorake/verisart-go/types"
	transferStatus "github.com/quinlanmorake/verisart-go/types/transferStatus"	

	database "github.com/quinlanmorake/verisart-go/database"	
)

func AcceptTransfer(user businessTypes.User, certificateId coreTypes.String) coreTypes.Result {
	var certificate businessTypes.Certificate

	// Firstly, the certificate must still exists, if doesn't update the transfer and return failure
	if _certificate, loadCertificateResult := GetCertificateById(certificateId); loadCertificateResult.IsNotOk() {
		if loadCertificateResult.Code == errorCodes.NO_CERTIFICATE_WITH_THE_GIVEN_ID_COULD_BE_FOUND {
			updatePotentialTransferRecordToReflectDeletedCertificate(certificateId)
		}
		return loadCertificateResult
	} else {
		certificate = _certificate
	}

	// Then validate the result, also loads the transfer
	if transfer, validateBeforeAcceptingResult := validateBeforeAcceptingTransfer(user, certificate); validateBeforeAcceptingResult.IsNotOk() {
		return validateBeforeAcceptingResult
	} else {
		// First transfer the certificate, then update the transfer result
		certificate.OwnerId = user.Id
		if updateCertificateOwnerResult := database.Update(&certificate); updateCertificateOwnerResult.IsNotOk() {
			return updateCertificateOwnerResult
		}

		// Now update transfer
		transfer.ActionedOn = coreTypes.GenerateTimestamp()
		transfer.Status = transferStatus.ACCEPTED
		if updateTransferResult := database.Update(&transfer); updateTransferResult.IsNotOk() {
			return updateTransferResult
		}

		// Finally update transfer state on certificate
		certificate.Transfer.Status = transferStatus.ACCEPTED
		return database.Update(&certificate)
	}
}

func validateBeforeAcceptingTransfer(user businessTypes.User, certificate businessTypes.Certificate) (businessTypes.Transfer, coreTypes.Result) {
	var transfer businessTypes.Transfer

	// Transfer must exist for the certificate	
	if _transfer, loadTransferForCertificateResult := loadTransferForCertifcate(certificate); loadTransferForCertificateResult.IsNotOk() {
		return businessTypes.NewEmptyTransfer(), loadTransferForCertificateResult
	} else {
		transfer = _transfer
	}
	
	// The user accepting the transfer, must have the same email address as that user the transfer is destined for
	if transfer.To.IsEqual(user.Email) == false {
		return transfer, coreTypes.NewResultFromErrorCode(errorCodes.INSUFFICIENT_PERMISSION_TO_ACCEPT_THE_TRANSFER)		
	}

	return transfer, coreTypes.NewSuccessResult()
}

func updatePotentialTransferRecordToReflectDeletedCertificate(certificateId coreTypes.String) {
	// One scenario that may cause this is an invalid certificate id, so we cannot assume any transfers exist
	// That is, if we can't find any, just return
	// Calle doesn't care about the result of this operation, callee assumes it went fine, so we don't return anything

	certificate := businessTypes.Certificate{
		Id: certificateId,
	}
	
	if transfer, loadTransferForCertificateResult := loadTransferForCertifcate(certificate); loadTransferForCertificateResult.IsNotOk() {
		// As noted, just ignore failure
		return
	} else {
		transfer.ActionedOn = coreTypes.GenerateTimestamp()
		transfer.Status = transferStatus.ANNULED

		database.Update(&transfer)
	}
}