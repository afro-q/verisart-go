package types

type UserId String

func (u UserId) Length() int {
	return String(u).Length()
}
