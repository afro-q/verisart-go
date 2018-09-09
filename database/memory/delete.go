package memory

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"
	
	dbTypes "github.com/quinlanmorake/verisart-go/database/types"	
)

func (m *MemoryDb) Delete(record dbTypes.DbRecord) coreTypes.Result {
	if (record.GetId().Length() == 0) {
		return coreTypes.NewResultFromErrorCode(errorCodes.DATABASE_DELETE_NO_ID_WAS_PROVIDED)
	}
	
	if checkDbResult := m.IsInitialized(); checkDbResult.IsNotOk() {
		return checkDbResult
	}

	transaction := m.Db.Txn(true)

	if deleteError := transaction.Delete(record.GetTableName().ToString(), record); deleteError != nil {
		transaction.Abort()

		return coreTypes.Result{
			Code: errorCodes.DATABASE_DELETE_FAILED, 
			Message: errorMessages.ErrorMessage(deleteError.Error()),
		}
	} else {
		transaction.Commit()
	}
	
	return coreTypes.NewSuccessResult()
}
