package handlers

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type httpRequest struct {
	Title coreTypes.String `json:"title"`
}
