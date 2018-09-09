package memory

import (
	"testing"
)

func Add_Works(db MemoryDb, t *testing.T) {
	newUser := UserTestingObject{
		Name: "Test User",
	}

	if addResult := db.Add(&newUser); addResult.Result.IsNotOk() {
		t.Error(addResult.Result.Error())
	}
}

func Test_Add_Fails_When_Db_Is_Not_Initialized(t *testing.T) {
}

func Test_Add_Fails_When_Record_Provided_Already_Has_An_Id(t *testing.T) {
}