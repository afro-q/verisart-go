package authentication

import (
	"net/http"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	
	businessTypes "github.com/quinlanmorake/verisart-go/types"
)

const JWT_KEY = "Authorization"

func GetUserFromTokenInHeaders(r *http.Request) (businessTypes.User, coreTypes.Result) {
	if tokenFromHeader := coreTypes.String(r.Header.Get(JWT_KEY)); tokenFromHeader.Length() == 0 {
		return businessTypes.User{}, coreTypes.NewResultFromErrorCode(errorCodes.JWT_AUTHORIZATION_HEADER_WAS_NOT_SET)
	} else {
		return ParseToken(tokenFromHeader)
	}	
}
