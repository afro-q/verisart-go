package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	transferStatus "github.com/quinlanmorake/verisart-go/types/transferStatus"		
)

type TransferState struct {
	CreatedAt coreTypes.IsoTimeString `json:"createdAt"`
	CreatedBy coreTypes.UserId        `json:"originator"`

	Status transferStatus.TransferStatus `json:"status"`
}
