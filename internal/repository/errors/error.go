package errors

type Code int64

func (c Code) Error() string {
	return "unknown error"
}
