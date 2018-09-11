package user

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	database "github.com/quinlanmorake/verisart-go/database"
)

/*
 This just populates the database with some users
*/
func Init() coreTypes.Result {
	for index, _ := range defaultUsers {
		if addResult := database.Add(&defaultUsers[index]); addResult.Result.IsNotOk() {
			return addResult.Result
		}
	}

	return coreTypes.NewSuccessResult()
}
