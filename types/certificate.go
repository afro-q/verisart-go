package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"	
)

type Certificate struct {
	CreatedAt coreTypes.DateString `json:"createdAt"`
	Id        coreTypes.String     `json:"id"`

	Note    coreTypes.String `json:"note"`
	OwnerId coreTypes.UserId `json:"ownerId"`

	Title    coreTypes.String `json:"title"`
	Transfer TransferState    `json:"transfer"`

	Year coreTypes.Year `json:"year"`
}

func (c *Certificate) GetId() coreTypes.String {
	return c.Id
}

func (c *Certificate) SetId(id coreTypes.String) {
	c.Id = id
}

func (c *Certificate) GetTableName() coreTypes.String {
	return coreTypes.String(tableNames.CERTIFICATES)
}

func NewEmptyCertificate() Certificate {
	return Certificate{}
}

func (c *Certificate) CopyEditableFields(copyFrom Certificate) {
	c.Note = copyFrom.Note
	c.Title = copyFrom.Title
	c.Year = copyFrom.Year
}