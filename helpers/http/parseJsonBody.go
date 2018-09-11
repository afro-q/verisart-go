package http

import (
	"encoding/json"
	"io"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"			
)

func ParseJsonBody(body io.ReadCloser, target interface{}) coreTypes.Result {
	defer func() {
		_ = body.Close()
	}()

	decoder := json.NewDecoder(body)
	if parseError := decoder.Decode(target); parseError != nil {
		return coreTypes.Result {
			Code: errorCodes.INVALID_REQUEST_BODY,
			Message: errorMessages.ErrorMessage(parseError.Error()),
		}
	} else {
		return coreTypes.NewSuccessResult()
	}
}
