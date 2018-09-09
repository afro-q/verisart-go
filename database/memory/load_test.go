package memory

import (
	"encoding/json"
	"testing"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"	
	
	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"	
)

func Load_Works(db MemoryDb, t *testing.T) {
	resultHandler := dbTypes.DataHandler(func (dbRows [][]byte) coreTypes.Result {
		// Check we have data
		if len(dbRows) == 0 {
			t.Error("No rows were loaded from the database")
			return coreTypes.Result {
				Code: errorCodes.UNIT_TEST_GENERIC_FAILURE,
			}
		}
		
		usersInDb := make([]UserTestingObject, len(dbRows))

		// Check the data can marshal
		for index, jsonRowData := range dbRows {
			usersInDb[index] = UserTestingObject{}

			if unmarshalError := json.Unmarshal(jsonRowData, &usersInDb[index]); unmarshalError != nil {
				t.Error(unmarshalError.Error())
				return coreTypes.Result{
					Code: errorCodes.UNIT_TEST_GENERIC_FAILURE,
				}
			}
		}

		// We won't validate data content, just check there is something there
		for _, user := range usersInDb {
			if user.Id.Length() == 0 {
				t.Error("A user was found with no id")
				return coreTypes.Result {
					Code: errorCodes.UNIT_TEST_GENERIC_FAILURE,
				}
			}
		}

		return coreTypes.NewSuccessResult()
	})

	if loadResult := db.Load(tableNames.USERS, resultHandler); loadResult.IsNotOk() {
		t.Error(loadResult.Error())
	}
}
