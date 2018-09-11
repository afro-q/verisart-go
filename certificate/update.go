package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	businessTypes "github.com/quinlanmorake/verisart-go/types"
	
	database "github.com/quinlanmorake/verisart-go/database"
)

/*
 Only validation is that the user updating the certicate is allowed to do so
*/
func Update(userId coreTypes.UserId, certificateId coreTypes.String, certificateWithNewData businessTypes.Certificate) (businessTypes.Certificate, coreTypes.Result) {
	var certificateToUpdate businessTypes.Certificate

	if _certificate, getCertificateByIdResult := GetCertificateById(certificateId); getCertificateByIdResult.IsNotOk() {
		return businessTypes.NewEmptyCertificate(), getCertificateByIdResult
	} else {
		certificateToUpdate = _certificate
	}

	if certificateToUpdate.OwnerId.IsEqual(userId) == false {
		return businessTypes.NewEmptyCertificate(), coreTypes.NewResultFromErrorCode(errorCodes.INSUFFICIENT_PERMISSIONS_TO_MODIFY_THE_REQUESTED_CERTIFICATE)
	}
	
	certificateToUpdate.CopyEditableFields(certificateWithNewData)

	if updateResult := database.Update(&certificateToUpdate); updateResult.IsNotOk() {
		return businessTypes.NewEmptyCertificate(), updateResult
	} else {
		return certificateToUpdate, coreTypes.NewSuccessResult()
	}
}