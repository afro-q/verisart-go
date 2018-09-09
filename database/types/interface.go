package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

// NOTE: Will likely need to change to accept config in the later
type Init interface {
	Init() coreTypes.Result
}

type Adder interface {
	Add(DbRecord) AddResult
}

type Updater interface {
	/*
	   We are not adding extra work by implementing a comparer and updating only required entries,
	   the currently logic is going to update the entire object with the provided one
	*/
	Update(DbRecord) coreTypes.Result
}

/*
 In this case we are not creating another type, as we are only implementing "delete 1"
 As such, if not exactly 1 is not deleted, there will be a failure
*/
type Deleter interface {
	Delete(DbRecord) coreTypes.Result
}

/*
  We load data by calling load, the running the closure giving it the provided data.
  Load as such can fail at 2 points, either by the database code, trying to load, or by the provided closure.

  More importantly, being that this is an in-memory store anyway, we don't bother with filtering within the db layer.
  When this moves to a db, a filtering mechansim will have to be added in.  
*/
type Loader interface {
	Load(string, DataHandler) coreTypes.Result
}

type Database interface {
	Init
	Adder
	Updater
	Deleter
	Loader
}
