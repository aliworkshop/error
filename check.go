package error

func IsNotFoundError(err ErrorModel) bool {
	return err.Type() == TypeNotFound
}
