package types

type UserId String

func (u UserId) Length() int {
	return String(u).Length()
}

func (u UserId) IsEqual(compareTo UserId) bool {
	return String(u).IsEqual(String(compareTo))
}