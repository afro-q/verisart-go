package memory

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

func (m *MemoryDb) Delete(rowId coreTypes.String) coreTypes.Result {
	return coreTypes.NewSuccessResult()
}
