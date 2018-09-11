package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	businessTypes "github.com/quinlanmorake/verisart-go/types"
)

func loadTransferForCertifcate(certificate businessTypes.Certificate) (businessTypes.Transfer, coreTypes.Result) {
	// Assuming the certificate record exists in the database as the data type passed in is the actual certificate

	if transfers, loadAllTransfersResult := LoadAllTransfers(); loadAllTransfersResult.IsNotOk() {
		return businessTypes.NewEmptyTransfer(), loadAllTransfersResult
	} else {
		for index, transfer := range transfers {
			if transfer.CertificateId.IsEqual(certificate.Id) {
				return transfers[index], coreTypes.NewSuccessResult()
			}
		}

		return businessTypes.NewEmptyTransfer(), coreTypes.NewResultFromErrorCode(errorCodes.NO_TRANSFER_FOR_THE_GIVEN_CERTIFICATE_ID_COULD_BE_FOUND)
	}
}