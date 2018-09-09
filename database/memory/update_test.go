package memory

import (
	"encoding/json"
	"testing"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	
	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"		
)

func Update_Works(db MemoryDb, t *testing.T) {
	// This tests runs after "Load_Works", so we know load works and that there's data in the database
	// and that the data is valid, as such, will use load to get an id

	userToEdit := UserTestingObject{}
	getUserToEdit := dbTypes.DataHandler(func (dbRows [][]byte) coreTypes.Result {
		json.Unmarshal(dbRows[0], &userToEdit)

		return coreTypes.NewSuccessResult()
	})
		
	if loadResult := db.Load(tableNames.USERS, getUserToEdit); loadResult.IsNotOk() {
		t.Error(loadResult.Error())
	}

	userToEdit.Name = "James"
	if updateResult := db.Update(&userToEdit); updateResult.IsNotOk() {
		t.Error(updateResult.Error())
	}
}
