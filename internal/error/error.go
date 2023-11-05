package error

type Code int64

const (
	Unknown Code = iota
	DeadlineExceeded
	Canceled
	Internal
	NotFound
	AlreadyExists
	BadRequest
	PermissionDenied
	Unauthenticated
	Unimplemented
	Unavailable
)
