package http

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	businessTypes "github.com/quinlanmorake/verisart-go/types"
	
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
)

type createRequest struct {
	Note    coreTypes.String `json:"note"`
	Title   coreTypes.String `json:"title"`	
	Year    coreTypes.Year   `json:"year"`	
}

func (c createRequest) Validate() coreTypes.Result {
	// These are made up business rules
	if c.Title.Length() == 0 {
		return coreTypes.NewResultFromErrorCode(errorCodes.CERTIFICATE_MISSING_TITLE)
	}
	
	if c.Year < 2000 {
		return coreTypes.NewResultFromErrorCode(errorCodes.CERTIFICATE_HAS_INVALID_YEAR)
	}

	return coreTypes.NewSuccessResult()
}

func (c createRequest) ToCertificate() businessTypes.Certificate {
	return businessTypes.Certificate {
		Note: c.Note,
		Title: c.Title,
		Year: c.Year,
	}
}