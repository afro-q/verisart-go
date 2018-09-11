package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"	
	transferStatus "github.com/quinlanmorake/verisart-go/types/transferStatus"		
)

type Transfer struct {
	To coreTypes.EmailAddress `json:"to"`
	Status transferStatus.TransferStatus `json:"status"`
}