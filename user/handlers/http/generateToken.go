package http

import (
	"net/http"

	gorillaMux "github.com/gorilla/mux"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"

	httpHelpers "github.com/quinlanmorake/verisart-go/helpers/http"

	user "github.com/quinlanmorake/verisart-go/user"
)

/*
 This generates a token that can be used to make any of the authenticated http requests
 It should be set as the value of the "Authorization" HTTP header
*/
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	response := generateTokenResponse{}
	defer func() {
		httpHelpers.WriteResponse(w, response)
	}()

	var userId coreTypes.UserId
	requestVariables := gorillaMux.Vars(r)
	if userIdParam, userIdFound := requestVariables["userId"]; userIdFound == false {
		response.Error = coreTypes.NewResultFromErrorCode(errorCodes.USER_ID_IS_INVALID)
	} else {
		userId = coreTypes.UserId(userIdParam)
	}

	if userId.Length() == 0 {
		response.Error = coreTypes.NewResultFromErrorCode(errorCodes.USER_ID_IS_INVALID)
		return
	}

	response.Jwt, response.Error = user.GenerateToken(userId)
}
