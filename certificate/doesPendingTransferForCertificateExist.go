package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	businessTypes "github.com/quinlanmorake/verisart-go/types"
	transferStatus "github.com/quinlanmorake/verisart-go/types/transferStatus"	
)

func doesPendingTransferForCertificateExist(certificate businessTypes.Certificate) (bool, coreTypes.Result) {
	// Assuming the certificate record exist in the database as the data type passed in is the actual certificate
	
	if transfers, loadAllTransfersResult := LoadAllTransfers(); loadAllTransfersResult.IsNotOk() {
		return false, loadAllTransfersResult
	} else {
		for _, transfer := range transfers {
			if transfer.CertificateId.IsEqual(certificate.Id) && (transfer.Status == transferStatus.CREATED) {
				return true, coreTypes.NewSuccessResult()
			}
		}

		return false, coreTypes.NewSuccessResult()		
	}	
}