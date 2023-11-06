package errors

type Code int64

const (
	Unknown Code = 0

	DontHaveAvailableRooms       Code = 1
	OrderForThatTimeAlreadyExist Code = 2
)

func (c Code) Error() string {
	switch c {
	case DontHaveAvailableRooms:
		return "don't have available rooms"
	case OrderForThatTimeAlreadyExist:
		return "order for that time already exist"
	default:
		return "unknown"
	}
}
