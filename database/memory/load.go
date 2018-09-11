package memory

/*
 As the library returns an empty interface, we're going to json marshal to bytes, and let the callee unmarshal
*/

import (
	"encoding/json"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"

	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
)

func (m *MemoryDb) Load(tableName string, dataHandler dbTypes.DataHandler) coreTypes.Result {
	if checkDbResult := m.IsInitialized(); checkDbResult.IsNotOk() {
		return checkDbResult
	}

	transaction := m.Db.Txn(false)
	defer transaction.Abort()

	// NOTE: We know that every table has an id index, we unit test for that.
	if rowIterator, loadError := transaction.Get(tableName, "id"); loadError != nil {
		return newFailureResult(loadError.Error())
	} else {
		dbRows := make([][]byte, 0)

		for dataRow := rowIterator.Next(); dataRow != nil; dataRow = rowIterator.Next() {
			if rowAsJsonByteArray, marshalError := json.Marshal(dataRow); marshalError != nil {
				return newFailureResult(marshalError.Error())
			} else {
				dbRows = append(dbRows, rowAsJsonByteArray)
			}
		}

		return dataHandler(dbRows)
	}
}

func newFailureResult(message string) coreTypes.Result {
	return coreTypes.Result{
		Code:    errorCodes.DATABASE_LOAD_FAILED,
		Message: errorMessages.ErrorMessage(message),
	}
}
