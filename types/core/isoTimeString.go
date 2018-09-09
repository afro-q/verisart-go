package types

type IsoTimeString string

func GenerateTimestamp() IsoTimeString {
	return IsoTimeString("TODO")
}

func (its IsoTimeString) ToDateString() DateString {
	return DateString("TODO")
}
