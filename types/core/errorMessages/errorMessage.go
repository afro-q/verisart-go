package errorMessages

/*
 The ErrorMessage is envisioned as a type especially because we have an API and are thinking about internationlization,
 but even without, something like a table of error messages in the database would suffice to show value of this.

 At the moment we are going to get the error message by looking up in a dictionary, but in the future we would like
 to get the error message from a table; given a language id, we could easily add i18n too.

 There is a descrepency at the moment in that we are sometimes returning the error message as given by an underlying
 library instead of our own.
 One can identify those codes by running "utilities/getErrorCodesWithNoMessages"
 When / Should business dictate that this is not the desired behaviour, if only for 1 error code, one must just add
 a message to the dictionary of error messages; and if for all behaviour, then one must modify the
  NewResultFromErrorCode function in the "types/result.go" file.

*/

type ErrorMessage string

func (em ErrorMessage) ToString() string {
	return string(em)
}
