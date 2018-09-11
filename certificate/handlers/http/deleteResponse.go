package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type deleteResponse struct {
	Error coreTypes.Result `json:"error"`
}
