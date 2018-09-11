package memory

import (
	"encoding/json"
	"testing"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
	tableNames "github.com/quinlanmorake/verisart-go/database/types/tableNames"
)

func Delete_Works(db MemoryDb, t *testing.T) {
	// Just like edit, This tests runs after "Load_Works", so we know load works and that there's
	// data in the database and that the data is valid, as such, will use load to get an id

	userToDelete := UserTestingObject{}
	getUserToDelete := dbTypes.DataHandler(func(dbRows [][]byte) coreTypes.Result {
		json.Unmarshal(dbRows[0], &userToDelete)

		return coreTypes.NewSuccessResult()
	})

	if loadResult := db.Load(tableNames.USERS, getUserToDelete); loadResult.IsNotOk() {
		t.Error(loadResult.Error())
	}

	if deleteResult := db.Delete(&userToDelete); deleteResult.IsNotOk() {
		t.Error(deleteResult.Error())
	}
}
