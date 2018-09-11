package authentication

import (
	"encoding/json"

	jose "github.com/dvsekhvalnov/jose2go"

	businessTypes "github.com/quinlanmorake/verisart-go/types"
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"
)

func ParseToken(token coreTypes.String) (businessTypes.User, coreTypes.Result) {
	if payload, _, decodeError := jose.Decode(token.ToString(), publicKey); decodeError != nil {
		return businessTypes.User{}, coreTypes.Result{
			Code:    errorCodes.JWT_COULD_NOT_DECODE_TOKEN,
			Message: errorMessages.ErrorMessage(decodeError.Error()),
		}
	} else {
		tokenData := TokenData{}

		if unmarshalError := json.Unmarshal([]byte(payload), &tokenData); unmarshalError != nil {
			return businessTypes.User{}, coreTypes.Result{
				Code:    errorCodes.JWT_COULD_NOT_UNMARSHAL_TOKEN,
				Message: errorMessages.ErrorMessage(unmarshalError.Error()),
			}
		}

		if validityCheck := tokenData.IsValid(); validityCheck.IsNotOk() {
			return businessTypes.User{}, validityCheck
		}

		return tokenData.User, coreTypes.NewSuccessResult()
	}
}
