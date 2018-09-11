package user

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	businessTypes "github.com/quinlanmorake/verisart-go/types"
)

func GetUserById(userId coreTypes.UserId) (businessTypes.User, coreTypes.Result) {
	var userToReturn businessTypes.User

	if users, loadAllUsersResult := LoadAllUsers(); loadAllUsersResult.IsNotOk() {
		return businessTypes.User{}, loadAllUsersResult
	} else {
		foundUserWithId := false

		for _, userInDb := range users {
			if userInDb.Id == userId {
				userToReturn = userInDb
				foundUserWithId = true
			}
		}
		
		if foundUserWithId == false {
			return businessTypes.User{}, coreTypes.NewResultFromErrorCode(errorCodes.NO_USER_WITH_THE_GIVEN_ID_EXISTS)
		}		
	}

	return userToReturn, coreTypes.NewSuccessResult()
}
