package database

import (
	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	config "github.com/quinlanmorake/verisart-go/config"

	memoryDb "github.com/quinlanmorake/verisart-go/database/memory"
)

var Db dbTypes.Database
var isDbInitialized bool = false

type dbConstructor func() dbTypes.Database

var dbConstructorMap map[string]dbConstructor = map[string]dbConstructor{
	"memory": dbConstructor(func() dbTypes.Database {
		// NOTE: Because init sets values on the object, we implement the interface on the pointer
		return &memoryDb.MemoryDb{}
	}),
}

// Let's only use the object in config package directly when we absolutely need to
func Init(appConfig config.Config) coreTypes.Result {
	if constructor, constructorFound := dbConstructorMap[appConfig.Database.ToLowercaseString()]; constructorFound == false {
		return coreTypes.NewResultFromErrorCode(errorCodes.INVALID_DATABASE_CONFIG)
	} else {
		Db = constructor()
		isDbInitialized = true

		return Db.Init()
	}
}

// We only have 1 db, so may as well add helpers
func Add(record dbTypes.DbRecord) dbTypes.AddResult {
	if dbCheck := validateInitialized(); dbCheck.IsNotOk() {
		return dbTypes.AddResult{
			Result: dbCheck,
		}
	} else {
		return Db.Add(record)
	}
}

type resultReturner func() coreTypes.Result

func Update(record dbTypes.DbRecord) coreTypes.Result {
	return validateAndCall(resultReturner(func() coreTypes.Result {
		return Db.Update(record)
	}))
}

func Delete(record dbTypes.DbRecord) coreTypes.Result {
	return validateAndCall(resultReturner(func() coreTypes.Result {
		return Db.Delete(record)
	}))
}

func Load(tableName string, dataHandler dbTypes.DataHandler) coreTypes.Result {
	return validateAndCall(resultReturner(func() coreTypes.Result {
		return Db.Load(tableName, dataHandler)
	}))
}

// Just done to showcase closure usage and anonymous functions
func validateAndCall(functionToExecuteIfValidated resultReturner) coreTypes.Result {
	if dbCheck := validateInitialized(); dbCheck.IsNotOk() {
		return dbCheck
	} else {
		return functionToExecuteIfValidated()
	}
}

func validateInitialized() coreTypes.Result {
	if isDbInitialized == false {
		return coreTypes.NewResultFromErrorCode(errorCodes.DATABASE_NOT_INITIALIZED)
	} else {
		return coreTypes.NewSuccessResult()
	}
}
