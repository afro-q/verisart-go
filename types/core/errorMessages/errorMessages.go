package errorMessages

import (
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
)

var ErrorMessages = map[errorCodes.ErrorCode]ErrorMessage{
	errorCodes.USER_ID_IS_INVALID:            ErrorMessage("The id on the provided user is not valid"),
	errorCodes.EMAIL_IS_INVALID:              ErrorMessage("The given email address is not a valid email address"),
	errorCodes.NO_USER_WITH_THE_EMAIL_EXISTS: ErrorMessage("No user with the given email address could be found"),

	errorCodes.INVALID_DATABASE_CONFIG:           ErrorMessage("The database element in the config valid is invalid, it is either empty, or is not a valid database type"),
	errorCodes.DATABASE_INIT_ERROR:               ErrorMessage("An error occured initializing the database"),
	errorCodes.DATABASE_NOT_INITIALIZED:          ErrorMessage("An operation is being attempted against a database that is not initialized"),
	errorCodes.DATABASE_ADD_DID_NOT_RETURN_AN_ID: ErrorMessage("On adding an object to the database, no id was generated it"),
}
