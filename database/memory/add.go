package memory

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"	
	
	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
)

func (m *MemoryDb) Add(entryToAdd dbTypes.DbRecord) dbTypes.AddResult {
	if checkDbResult := m.IsInitialized(); checkDbResult.IsNotOk() {
		return dbTypes.NewAddResult(checkDbResult)
	}

	// The object should not have an id
	if (entryToAdd.GetId().Length() > 0) {
		return dbTypes.NewAddResult(coreTypes.NewResultFromErrorCode(errorCodes.DATABASE_ADD_OBJECT_ALREADY_HAS_AN_ID))
	}
	entryToAdd.SetId(coreTypes.String(NewId()))
	
	transaction := m.Db.Txn(true)

	if addError := transaction.Insert(entryToAdd.GetTableName().ToString(), entryToAdd); addError != nil {
		transaction.Abort()

		result := coreTypes.Result{
			Code: errorCodes.DATABASE_ADD_FAILED,
			Message: errorMessages.ErrorMessage(addError.Error()),
		}

		return dbTypes.NewAddResult(result)
	} else {
		transaction.Commit()
	}

	// We are only setting id for the purposes of the library, as such this is redundant, but it may be the case
	// that the callee does not hold onto the object, expected to attain the its id from here
	return dbTypes.AddResult{
		Result: coreTypes.NewSuccessResult(),
		NewEntryId: entryToAdd.GetId(),
	}
}
