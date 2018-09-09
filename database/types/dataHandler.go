package types

import (
	"encoding/json"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type DataHandler func([]json.Unmarshaler) coreTypes.Result
