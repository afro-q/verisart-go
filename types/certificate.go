package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type Certificate struct {
	CreatedAt coreTypes.DateString `json:"createdAt"`
	Id        coreTypes.String     `json:"id"`

	Note    coreTypes.String `json:"note"`
	OwnerId coreTypes.UserId `json:"ownerId"`

	Title    coreTypes.String `json:"title"`
	Transfer TransferState    `json:"transfer"`

	Year coreTypes.Year `json:"year"`
}
