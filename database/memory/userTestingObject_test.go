package memory

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"
)

type UserTestingObject struct {
	Id   coreTypes.String `json:"id"`
	Name coreTypes.String `json:"name"`
}

func (u UserTestingObject) GetId() coreTypes.String {
	return u.Id
}

func (u *UserTestingObject) SetId(id coreTypes.String) {
	u.Id = id
}

func (u UserTestingObject) GetTableName() coreTypes.String {
	return coreTypes.String(tableNames.USERS)
}
