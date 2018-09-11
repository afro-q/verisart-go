package handlers

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type httpRequest struct {
	Title   coreTypes.String `json:"title"`
	OwnerId coreTypes.UserId `json:"ownerId"`
	Year    coreTypes.Year   `json:"year"`
	Note    coreTypes.String `json:"note"`
}
