package types

/*
 I do not like that we end up have to write the same word twice "errorCodes.ErrorCode", but this is the only place
 we have to, so its only in this file; and then it provides clarity in using the constants everywhere else; so I feel
 it an okay tradeoff.
*/

import (
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"
)

type Result struct {
	Code    errorCodes.ErrorCode       `json:"code"`
	Message errorMessages.ErrorMessage `json:"message"`
}

// Implement the error interface
func (r Result) Error() string {
	return string(r.Message)
}

func NewSuccessResult() Result {
	return Result{
		Code: errorCodes.ErrorCode(0),
	}
}

func NewResultFromErrorCode(errorCode errorCodes.ErrorCode) Result {
	// By design, we will just return an empty message if there is no corresponding message for the error code
	return Result{
		Code:    errorCode,
		Message: errorMessages.ErrorMessages[errorCode],
	}
}

// I just find this clearer, would obviously need agreement with dev team for consistency sake
func (r Result) IsNotOk() bool {
	return r.Code != errorCodes.ErrorCode(0)
}
