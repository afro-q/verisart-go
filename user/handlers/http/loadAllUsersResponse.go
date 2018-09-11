package http

import (
	businessTypes "github.com/quinlanmorake/verisart-go/types"
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type loadAllUsersResponse struct {
	Error coreTypes.Result     `json:"error"`
	Users []businessTypes.User `json:"users"`
}
