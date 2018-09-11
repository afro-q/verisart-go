package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	
	businessTypes "github.com/quinlanmorake/verisart-go/types"	
)

type updateRequest struct {
	Note coreTypes.String `json:"note"`
	Title coreTypes.String `json:"title"`
	Year coreTypes.Year `json:"year"`
}

func (u updateRequest) ToCertificate() businessTypes.Certificate {
	return businessTypes.Certificate {
		Note: u.Note,
		Title: u.Title,
		Year: u.Year,
	}
}