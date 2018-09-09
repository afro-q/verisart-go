package handlers

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type httpResponse struct {
	Error       coreTypes.Result      `json:"error"`
}
