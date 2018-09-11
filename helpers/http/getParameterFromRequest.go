package http

import (
	"net/http"
	
	gorillaMux "github.com/gorilla/mux"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"	
)

func GetParameterFromRequest(r *http.Request, parameterName string, errorCodeIfParameterIsNotFound errorCodes.ErrorCode) (parameterValue coreTypes.String, result coreTypes.Result) {
	requestVariables := gorillaMux.Vars(r)
	if _requestParameter, requestParameterFound := requestVariables[parameterName]; requestParameterFound == false {
		result = coreTypes.NewResultFromErrorCode(errorCodeIfParameterIsNotFound)
		return
	} else {
		parameterValue = coreTypes.String(_requestParameter)
	}

	if parameterValue.Length() == 0 {
		result = coreTypes.NewResultFromErrorCode(errorCodes.USER_ID_IS_INVALID)
	}

	return
}