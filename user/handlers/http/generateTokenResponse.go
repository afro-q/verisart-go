package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type generateTokenResponse struct {
	Error coreTypes.Result `json:"error"`
	Jwt   coreTypes.String `json:"jwt"`
}
