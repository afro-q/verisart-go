package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type TransferState struct {
	CreatedAt coreTypes.IsoTimeString `json:"createdAt"`
	CreatedBy coreTypes.UserId        `json:"originator"`

	Status TransferStatus `json:"status"`
}
