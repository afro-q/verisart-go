package types

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type DataHandler func([][]byte) coreTypes.Result
