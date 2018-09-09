package memory

import (
	"encoding/json"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
)

func (m *MemoryDb) Add(entryToAdd json.Marshaler) dbTypes.AddResult {
	if checkDbResult := m.IsInitialized(); checkDbResult.IsNotOk() {
		return dbTypes.AddResult{
			Result: checkDbResult,
		}
	}

	transaction := m.Db.Txn(true)

	transaction.Commit()
	
	return dbTypes.AddResult{
		Result: coreTypes.NewSuccessResult(),
	}
}
