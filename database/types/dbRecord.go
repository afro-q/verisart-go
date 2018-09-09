package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

/*
 This represents the minal required functionality that an object must implement
 so as to be able to be persisted.

 It is currently model strictly around the memory database library being used at the
 moment, and the model would change depending on additional libraries and their 
 requirements
*/

type DbRecord interface {	
	GetId() coreTypes.String	

	SetId(coreTypes.String)
	
	GetTableName() coreTypes.String
}