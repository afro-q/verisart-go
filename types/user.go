package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"	
)

type User struct {
	Id    coreTypes.UserId       `json:"id"`
	Email coreTypes.EmailAddress `json:"email"`
	Name  coreTypes.String       `json:"name"`
}

/*
 Not keeping a "createdAt" field here as doesn't seem relevant
*/
