package errors

type Code int64

const (
	Unknown          Code = 0
	DeadlineExceeded Code = 1
	Canceled         Code = 2
	Internal         Code = 3
	NotFound         Code = 4
	AlreadyExists    Code = 5
	BadRequest       Code = 6
	PermissionDenied Code = 7
	Unauthenticated  Code = 8
	Unimplemented    Code = 9
	Unavailable      Code = 10
)

func (c Code) Error() string {
	switch c {
	case DeadlineExceeded:
		return "deadline exceeded"
	case Canceled:
		return "canceled"
	case Internal:
		return "internal error"
	case NotFound:
		return "not found"
	case AlreadyExists:
		return "already exists"
	case BadRequest:
		return "bad request"
	case PermissionDenied:
		return "permission denied"
	case Unauthenticated:
		return "unauthenticated"
	case Unimplemented:
		return "unimplemented"
	case Unavailable:
		return "unavailable"
	default:
		return "unknown"
	}
}
