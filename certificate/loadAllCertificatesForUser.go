package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	businessTypes "github.com/quinlanmorake/verisart-go/types"	
)

func LoadAllCertificatesForUser(userId coreTypes.UserId) ([]businessTypes.Certificate, coreTypes.Result) {
	certificatesForUser := make([]businessTypes.Certificate, 0)
	
	if allCertificates, loadAllCertificatesResult := LoadAllCertificates(); loadAllCertificatesResult.IsNotOk() {
		return certificatesForUser, loadAllCertificatesResult
	} else {		
		for _, certificate := range allCertificates {
			if certificate.OwnerId.IsEqual(userId) {
				certificatesForUser = append(certificatesForUser, certificate)
			}
		}
	}

	return certificatesForUser, coreTypes.NewSuccessResult()
}