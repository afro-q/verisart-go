package types

type EmailAddress String

func (e EmailAddress) IsEqual(compareTo EmailAddress) bool {
	return String(e).IsEqual(String(compareTo))
}