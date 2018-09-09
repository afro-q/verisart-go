package database

import (
	dbTypes "github.com/quinlanmorake/verisart-go/database/types"
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	config "github.com/quinlanmorake/verisart-go/config"

	memoryDb "github.com/quinlanmorake/verisart-go/database/memory"
)

var Db dbTypes.Database

type dbConstructor func () dbTypes.Database

var dbConstructorMap map[string]dbConstructor = map[string]dbConstructor {
	"memory": dbConstructor(func () dbTypes.Database {
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
		return Db.Init()
	}
}
