package types

import (
	"time"
)

type DateString string

const (
	dateStringFormat = "Mon, 02 Jan 2006"
)
	
func NewDateStringTimestamp() DateString {
	return DateString(time.Now().Format(dateStringFormat))
}