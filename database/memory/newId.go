package memory

/*
 This database library does not generate ids, so we generate them here
*/

import (
	"github.com/satori/go.uuid"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

/*
 Leaning on the library once again, for all functionality, so no particular unit tests required
*/
func NewId() coreTypes.String {
	return coreTypes.String(uuid.NewV4().String())
}
