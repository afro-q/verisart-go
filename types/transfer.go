package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"		
	transferStatus "github.com/quinlanmorake/verisart-go/types/transferStatus"	
)

/*
 Even though only a subset of these properties are returned to the callee through the api,
 we don't let the requirements of the callee dictate the data object at this level.

 As needs be, we will create an http type / anonymous type / response particular type to
 abstract the api contract, instantiating that type from this one / using a factory to return it.
*/

type Transfer struct {
	ActionedOn coreTypes.IsoTimeString `json:"actionedOn"`

	CertificateId coreTypes.String `json:"certificateId"`
	CreatedAt coreTypes.IsoTimeString `json:"createdAt"`
	CreatedBy coreTypes.UserId        `json:"createdBy"`

	Id coreTypes.String `json:"id"`

	Status transferStatus.TransferStatus `json:"status"`

	To coreTypes.EmailAddress `json:"to"`
}

func NewTransfer(userId coreTypes.UserId, certificateId coreTypes.String, to coreTypes.EmailAddress) Transfer {
	return Transfer{
		CertificateId: certificateId,
		CreatedAt: coreTypes.GenerateTimestamp(),
		CreatedBy: userId,
		Status: transferStatus.CREATED,
		To: to,
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

func NewEmptyTransfer() Transfer {
	return Transfer{}
}

func (t *Transfer) GetId() coreTypes.String {
	return t.Id
}

func (t *Transfer) SetId(id coreTypes.String) {
	t.Id = id
}

func (t *Transfer) GetTableName() coreTypes.String {
	return coreTypes.String(tableNames.TRANSFERS)
}
