package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	
	database "github.com/quinlanmorake/verisart-go/database"
)

/*
 Only validation is that the user deleting the certicate is allowed to do so
*/
func Delete(userId coreTypes.UserId, certificateId coreTypes.String) coreTypes.Result {
	if certificate, getCertificateByIdResult := GetCertificateById(certificateId); getCertificateByIdResult.IsNotOk() {
		return getCertificateByIdResult
	} else {
		if certificate.OwnerId.IsEqual(userId) == false {
			return coreTypes.NewResultFromErrorCode(errorCodes.INSUFFICIENT_PERMISSIONS_TO_MODIFY_THE_REQUESTED_CERTIFICATE)
		}

		return database.Delete(&certificate)
	}
}