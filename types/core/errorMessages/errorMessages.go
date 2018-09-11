package errorMessages

import (
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
)

var ErrorMessages = map[errorCodes.ErrorCode]ErrorMessage{
	errorCodes.JWT_TOKEN_HAS_INVALID_ISSUER:         ErrorMessage("The token has an invalid issuer"),
	errorCodes.JWT_TOKEN_HAS_EXPIRED:                ErrorMessage("The JWT token has expired"),
	errorCodes.JWT_AUTHORIZATION_HEADER_WAS_NOT_SET: ErrorMessage("No value was provided in the Authorization header"),

	errorCodes.USER_ID_IS_INVALID:               ErrorMessage("The id on the provided user is not valid"),
	errorCodes.NO_USER_WITH_THE_GIVEN_ID_EXISTS: ErrorMessage("A user with the given id could not be found"),
	errorCodes.EMAIL_IS_INVALID:                 ErrorMessage("The given email address is not a valid email address"),
	errorCodes.NO_USER_WITH_THE_EMAIL_EXISTS:    ErrorMessage("No user with the given email address could be found"),

	errorCodes.INVALID_DATABASE_CONFIG:  ErrorMessage("The database element in the config valid is invalid, it is either empty, or is not a valid database type"),
	errorCodes.DATABASE_INIT_ERROR:      ErrorMessage("An error occured initializing the database"),
	errorCodes.DATABASE_NOT_INITIALIZED: ErrorMessage("An operation is being attempted against a database that is not initialized"),

	errorCodes.DATABASE_ADD_OBJECT_ALREADY_HAS_AN_ID: ErrorMessage("The record provided to be added to the database already has an id"),
	errorCodes.DATABASE_ADD_DID_NOT_RETURN_AN_ID:     ErrorMessage("On adding an object to the database, no id was generated it"),

	errorCodes.DATABASE_UPDATE_NO_ID_WAS_PROVIDED: ErrorMessage("An empty id was provided to the update database method"),

	errorCodes.ERROR_UNMARSHALING_USER_RECORD_FROM_DATABASE: ErrorMessage("An error occurred attempting to load a user reocrd from the database"),

	errorCodes.CERTIFICATE_MISSING_TITLE: ErrorMessage("No title was provided when creating or updating the certificate"),
	errorCodes.CERTIFICATE_HAS_INVALID_YEAR: ErrorMessage("The year on the certificate must have a value greater that '2000'"),
	errorCodes.ERROR_UNMARSHALING_CERTIFICATE_RECORD_FROM_DATABASE: ErrorMessage("An error occurred attempting to load a certificate record from the database"),
	errorCodes.NO_CERTIFICATE_WITH_THE_GIVEN_ID_COULD_BE_FOUND: ErrorMessage("The provided id does not match any certificate in the database"),
	errorCodes.INSUFFICIENT_PERMISSIONS_TO_MODIFY_THE_REQUESTED_CERTIFICATE: ErrorMessage("You do not have enough permission to modify the certicate - it belongs to another user"),
	errorCodes.CERTIFICATE_ID_IS_INVALID: ErrorMessage("An invalid id was provided as the certificate id parameter"),

	errorCodes.CERTIFICATE_ALREADY_HAS_PENDING_TRANSFER: ErrorMessage("The certificate already has a pending transfer"),
	errorCodes.ERROR_UNMARSHALING_TRANSFER_RECORD_FROM_DATABASE: ErrorMessage("An error occured attempting to load a transfer record from the database"),
	errorCodes.NO_TRANSFER_FOR_THE_GIVEN_CERTIFICATE_ID_COULD_BE_FOUND: ErrorMessage("No transfer corresponding to the certificate id could be found"),
	errorCodes.UNABLE_TO_TRANSFER_TO_YOURSELF: ErrorMessage("The receiver of the transfer must be different to the sender"),
	errorCodes.INSUFFICIENT_PERMISSION_TO_ACCEPT_THE_TRANSFER: ErrorMessage("You do not have enough permission to accept the certificate - it is destined for a user with a different email address"),
}
