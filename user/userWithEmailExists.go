package user

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"	
)

func UserWithEmailExists(emailAddress coreTypes.EmailAddress) (bool, coreTypes.Result) {
	if users, loadAllUsersResult := LoadAllUsers(); loadAllUsersResult.IsNotOk() {
		return false, loadAllUsersResult
	} else {
		for _, user := range users {
			if user.Email.IsEqual(emailAddress) {
				return true, coreTypes.NewSuccessResult()
			}
		}

		return false, coreTypes.NewSuccessResult()
	}
}