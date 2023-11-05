package errors

type Code int64

const (
	Unknown Code = iota

	DontHaveAvailableRooms
	OrderForThatTimeAlreadyExist
)

func (c Code) Error() string {
	switch c {
	case DontHaveAvailableRooms:
		return "don't have available rooms"
	default:
		return "unknown"
	}
}
