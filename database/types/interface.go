package types

/*
 We are going to lean on the Go Marhsaller interface for the implementation of database applicable types.

 In the future, as applicable, one could develop one's own interface, implemented by types able to be saved to / read from, the database
*/

import (
	"encoding/json"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

// NOTE: Will likely need to change to accept config in the later
type Init interface {
	Init() coreTypes.Result
}

type Adder interface {
	Add(json.Marshaler) AddResult
}

type Updater interface {
	/*
	   We are not adding extra work by implementing a comparer and updating only required entries,
	   the currently logic is going to update the entire object with the provided one
	*/
	Update(coreTypes.String, json.Marshaler) UpdateResult
}

/*
 In this case we are not creating another type, as we are only implementing "delete 1"
 As such, if not exactly 1 is not deleted, there will be a failure
*/
type Deleter interface {
	Delete(coreTypes.String) coreTypes.Result
}

/*
  We load data by calling load, the running the closure giving it the provided data.
  Load as such can fail at 2 points, either by the database code, trying to load, or by the provided closure.
*/
type Loader interface {
	Load(DataHandler) coreTypes.Result
}

type Database interface {
	Init
	Adder
	Updater
	Deleter
	Loader
}
