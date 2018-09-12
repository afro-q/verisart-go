package types

import (
	"time"
)

type IsoTimeString string

func GenerateTimestamp() IsoTimeString {
	return IsoTimeString(time.Now().Format(time.RFC3339))
}

func (its IsoTimeString) ToDateString() DateString {
	dateFromIsoTimeString, _ := time.Parse(time.RFC3339, string(its))
	return NewDateStringFromTime(dateFromIsoTimeString)
}
