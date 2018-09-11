package memory

/*
 Because we first need to add data in order to test that load works, or update,
 we just run those success tests as a suite
*/

import (
	"testing"
)

func Test_CRUD(t *testing.T) {
	db := MemoryDb{}

	if initResult := db.Init(); initResult.IsNotOk() {
		t.Error(initResult.Error())
		return
	}

	t.Run("Test_Add_Works", func(subT *testing.T) { Add_Works(db, subT) })
	t.Run("Test_Add_Works", func(subT *testing.T) { Add_Works(db, subT) })
	t.Run("Test_Load_Works", func(subT *testing.T) { Load_Works(db, subT) })
	t.Run("Test_Update_Works", func(subT *testing.T) { Update_Works(db, subT) })
	t.Run("Test_Delete_Works", func(subT *testing.T) { Delete_Works(db, subT) })
}
