package user

import (
	"time"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	businessTypes "github.com/quinlanmorake/verisart-go/types"

	authentication "github.com/quinlanmorake/verisart-go/authentication"
)

var DEFAULT_TOKEN_LIFETIME, _ = time.ParseDuration("20m")

func GenerateToken(userId coreTypes.UserId) (coreTypes.String, coreTypes.Result) {
	var userToGenerateTokenFor businessTypes.User

	if users, loadAllUsersResult := LoadAllUsers(); loadAllUsersResult.IsNotOk() {
		return coreTypes.NewEmptyString(), loadAllUsersResult
	} else {
		foundUserWithId := false

		for _, user := range users {
			if user.Id == userId {
				userToGenerateTokenFor = user
				foundUserWithId = true
			}
		}

		if foundUserWithId == false {
			return coreTypes.NewEmptyString(), coreTypes.NewResultFromErrorCode(errorCodes.NO_USER_WITH_THE_GIVEN_ID_EXISTS)
		}
	}

	return authentication.GenerateToken(userToGenerateTokenFor, DEFAULT_TOKEN_LIFETIME)
}
