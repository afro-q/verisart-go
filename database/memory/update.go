package memory

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"

	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
)

/*
 Note, no validation is happening on the object here; as such, if one would like to enforce that,
 add a validation member on the interface
*/

func (m *MemoryDb) Update(record dbTypes.DbRecord) coreTypes.Result {
	if record.GetId().Length() == 0 {
		return coreTypes.NewResultFromErrorCode(errorCodes.DATABASE_UPDATE_NO_ID_WAS_PROVIDED)
	}

	if checkDbResult := m.IsInitialized(); checkDbResult.IsNotOk() {
		return checkDbResult
	}

	transaction := m.Db.Txn(true)

	if updateError := transaction.Insert(record.GetTableName().ToString(), record); updateError != nil {
		transaction.Abort()

		return coreTypes.Result{
			Code:    errorCodes.DATABASE_UPDATE_FAILED,
			Message: errorMessages.ErrorMessage(updateError.Error()),
		}
	} else {
		transaction.Commit()
	}

	return coreTypes.NewSuccessResult()
}
