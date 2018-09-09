package memory

import (
	goMemDb "github.com/hashicorp/go-memdb"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"
)

type MemoryDb struct {
	Db *goMemDb.MemDB
}

func (m *MemoryDb) Init() coreTypes.Result {
	if db, err := goMemDb.NewMemDB(&dbSchema); err != nil {
		return coreTypes.Result{
			Code:    errorCodes.DATABASE_INIT_ERROR,
			Message: errorMessages.ErrorMessage(err.Error()),
		}
	} else {
		m.Db = db

		return coreTypes.NewSuccessResult()
	}
}

// Must always check we are initialized before attempting an operation
// In another database type, this would well be "IsConnected, CanConnect, etc
func (m MemoryDb) IsInitialized() coreTypes.Result {
	if m.Db == nil {
		return coreTypes.NewResultFromErrorCode(errorCodes.DATABASE_NOT_INITIALIZED)
	}

	return coreTypes.NewSuccessResult()
}
