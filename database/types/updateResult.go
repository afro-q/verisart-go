package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type UpdateResult struct {
	Result      coreTypes.Result `json:"error"`
	RowsUpdated int              `json:"rowsUpdated"`
}

/*
 We are not going to dictate business logic on if 0 rows updated means there was an error,
 So for the moment, the logic must be implemented in the consumer of the function if it is intended
*/
