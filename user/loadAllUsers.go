package user

import (
	"encoding/json"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
	businessTypes "github.com/quinlanmorake/verisart-go/types"

	database "github.com/quinlanmorake/verisart-go/database"
	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"
)

func LoadAllUsers() ([]businessTypes.User, coreTypes.Result) {
	var users []businessTypes.User

	databaseLoadHandler := dbTypes.DataHandler(func(dbRows [][]byte) coreTypes.Result {
		users = make([]businessTypes.User, len(dbRows))

		for index, jsonRowData := range dbRows {
			users[index] = businessTypes.User{}

			if unmarshalError := json.Unmarshal(jsonRowData, &users[index]); unmarshalError != nil {
				return coreTypes.NewResultFromErrorCode(errorCodes.ERROR_UNMARSHALING_USER_RECORD_FROM_DATABASE)
			}
		}

		return coreTypes.NewSuccessResult()
	})
	loadAllResult := database.Load(tableNames.USERS, databaseLoadHandler)

	return users, loadAllResult
}
