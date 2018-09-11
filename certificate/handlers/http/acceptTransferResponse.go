package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type acceptTransferResponse struct {
	Error coreTypes.Result `json:"error"`
}