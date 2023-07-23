package error

func IsNotFoundError(err error) bool {
	e, ok := err.(NotfoundError)
	if !ok {
		return false
	}
	return e.Type() == TypeNotFound
}
