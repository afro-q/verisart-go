package user

import (
	"time"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	authentication "github.com/quinlanmorake/verisart-go/authentication"
)

var DEFAULT_TOKEN_LIFETIME, _ = time.ParseDuration("20m")

func GenerateToken(userId coreTypes.UserId) (coreTypes.String, coreTypes.Result) {
	if userToGenerateTokenFor, getUserResult := GetUserById(userId); getUserResult.IsNotOk() {
		return coreTypes.NewEmptyString(), getUserResult
	} else {
		return authentication.GenerateToken(userToGenerateTokenFor, DEFAULT_TOKEN_LIFETIME)
	}
}
