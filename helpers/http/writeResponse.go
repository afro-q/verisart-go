package http

import (
	"encoding/json"
	"io"
)

func WriteResponse(writer io.Writer, responseData interface{}) {
	jsonWriter := json.NewEncoder(writer)

	// TODO: Handle an error here if one occurs
	jsonWriter.Encode(responseData)
}
