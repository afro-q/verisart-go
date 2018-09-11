package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type createTransferResponse struct {
	Error coreTypes.Result `json:"error"`
	Transfer Transfer `json:"transfer"`
}