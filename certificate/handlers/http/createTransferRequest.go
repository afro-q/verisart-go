package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type createTransferRequest struct {
	To coreTypes.EmailAddress `json:"to"`
}