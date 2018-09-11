package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	businessTypes "github.com/quinlanmorake/verisart-go/types"	
)

type loadCertificatesForUserResponse struct {
	Error coreTypes.Result `json:"error"`
	Certificates []businessTypes.Certificate `json:"certificate"`
}
