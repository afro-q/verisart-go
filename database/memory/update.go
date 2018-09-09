package memory

import (
	"encoding/json"

	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

func (m *MemoryDb) Update(rowId coreTypes.String, newEntry json.Marshaler) dbTypes.UpdateResult {
	return dbTypes.UpdateResult{
		Result: coreTypes.NewSuccessResult(),
	}
}
