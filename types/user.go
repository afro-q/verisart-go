package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"		
)

type User struct {
	Id    coreTypes.UserId       `json:"id"`
	Email coreTypes.EmailAddress `json:"email"`
	Name  coreTypes.String       `json:"name"`
}

/*
 Not keeping a "createdAt" field here as doesn't seem relevant
*/

func (u *User) GetId() coreTypes.String {
	return coreTypes.String(u.Id)
}

func (u *User) SetId(id coreTypes.String) {
	u.Id = coreTypes.UserId(id)
}

func (u *User) GetTableName() coreTypes.String {
	return coreTypes.String(tableNames.USERS)
}