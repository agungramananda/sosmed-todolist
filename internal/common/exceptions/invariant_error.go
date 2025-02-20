package exceptions

type InvariantError struct {
	Message string
}

func (e InvariantError) Error() string {
	return e.Message
}

func NewInvariantError(msg string) InvariantError {
	return InvariantError{Message: msg}
}