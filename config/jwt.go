package config

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"	
)

type jwt struct {
	Issuer coreTypes.String `json:"issuer"`
}