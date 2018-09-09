package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
)

type AddResult struct {
	Result     coreTypes.Result `json:"error"`
	NewEntryId coreTypes.String `json:"id"`
}

func NewAddResult(result coreTypes.Result) AddResult {
	return AddResult {
		Result: result,
	}
}

/*
 Whether or not an add operation succeeded depends on whether or not there was an error,
 and whether or not we have a valid id for the element added.

 Herein, we lean on GetError so as to determine whether or not we've succeed; the unit test does not couple that though; one can implement
 whatever one likes and as long as the logic is valid, the unit test will pass.
*/
func (ar AddResult) Succeeded() bool {
	// Just being clear
	if ar.GetError().IsNotOk() {
		return false
	} else {
		return true
	}
}

func (ar AddResult) GetError() coreTypes.Result {
	if ar.Result.IsNotOk() {
		return ar.Result
	}

	if ar.NewEntryId.Length() == 0 {
		return coreTypes.NewResultFromErrorCode(errorCodes.DATABASE_ADD_DID_NOT_RETURN_AN_ID)
	}

	return coreTypes.NewSuccessResult()
}
