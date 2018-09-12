package types

import (
	"time"
)

type DateString string

const (
	dateStringFormat = "Mon, 02 Jan 2006"
)
	
func NewDateStringTimestamp() DateString {
	return NewDateStringFromTime(time.Now())
}

func NewDateStringFromTime(t time.Time) DateString {
	return DateString(t.Format(dateStringFormat))
}