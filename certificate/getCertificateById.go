package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	businessTypes "github.com/quinlanmorake/verisart-go/types"	
)

func GetCertificateById(certificateId coreTypes.String) (businessTypes.Certificate, coreTypes.Result) {
	if certificates, loadAllCertificatesResult := LoadAllCertificates(); loadAllCertificatesResult.IsNotOk() {
		return businessTypes.NewEmptyCertificate(), loadAllCertificatesResult
	} else {
		for index, certificate := range certificates {
			if certificate.Id.IsEqual(certificateId) {
				return certificates[index], coreTypes.NewSuccessResult()
			}
		}

		return businessTypes.NewEmptyCertificate(), coreTypes.NewResultFromErrorCode(errorCodes.NO_CERTIFICATE_WITH_THE_GIVEN_ID_COULD_BE_FOUND)
	}
} 