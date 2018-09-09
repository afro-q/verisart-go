package memory

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
)

func (m *MemoryDb) Load(dataHandler dbTypes.DataHandler) coreTypes.Result {
	return coreTypes.NewSuccessResult()
}
