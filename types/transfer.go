package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
)

/*
 Even though only a subset of these properties are returned to the callee through the api,
 we don't let the requirements of the callee dictate the data object at this level.

 As needs be, we will create an http type / anonymous type / response particular type to
 abstract the api contract, instantiating that type from this one / using a factory to return it.
*/

type Transfer struct {
	ActionedOn coreTypes.IsoTimeString `json:"actionedOn"`

	CreatedAt coreTypes.IsoTimeString `json:"createdAt"`
	CreatedBy coreTypes.UserId        `json:"createdBy"`

	Id coreTypes.String `json:"id"`

	Status TransferStatus `json:"status"`

	To coreTypes.EmailAddress `json:"to"`
}

func NewTransfer() Transfer {
	return Transfer{
		CreatedAt: coreTypes.GenerateTimestamp(),
	}
}

/*
 The data will be stored in a database, and we don't know what will happen to it while
 it is there; as such, we have validity checks to ensure data integrity.

 That is, we know that requests will be validated, but we don't know if anything will
 happen while the data is in the database.
*/
func (t Transfer) IsValid() coreTypes.Result {
	if len(t.CreatedAt) == 0 {
		return coreTypes.NewResultFromErrorCode(errorCodes.TRANSFER_OBJECT_HAS_NO_CREATED_AT_VALUE)
	}

	return coreTypes.NewSuccessResult()
}
