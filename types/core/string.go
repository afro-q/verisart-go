package types

/*
 We wrap the Go string into on, adding additional helper methods
*/

import (
	"encoding/json"
	"strings"
)

type String string

// Just to make code easier by not adding an extra pair of surrounding brackets
func (s String) ToString() string {
	return strings.TrimSpace(string(s))
}

func (s String) ToLowercaseString() string {
	return strings.ToLower(s.ToString())
}

func (s String) Length() int {
	return len(strings.TrimSpace(s.ToString()))
}

func (s String) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ToString())
}

func (s *String) UnmarshalJSON(data []byte) error {
	asString := ""
	
	if unmarshalError := json.Unmarshal(data, &asString); unmarshalError != nil {
		return unmarshalError
	}
	
	*s = String(strings.TrimSpace(asString))
	
	return nil
}
