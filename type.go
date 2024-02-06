package error

import (
	"errors"
)

const (
	TypeValidation      ErrorType = "VALIDATION"
	TypeNotFound                  = "NOT_FOUND"
	TypeUnAuthorized              = "UNAUTHORIZED"
	TypeForbidden                 = "FORBIDDEN"
	TypeInternal                  = "INTERNAL"
	TypeDuplicate                 = "DUPLICATE"
	TypeTooManyRequests           = "TOO_MANY_REQUESTS"
)

type ErrorType string

func (error *err) Is(errType ErrorType) bool {
	return error.Type() == errType
}

func Is(err error, errType any) bool {
	if e, ok := err.(ErrorModel); ok {
		return e.Is(errType.(ErrorType))
	}

	return errors.Is(err, errType.(error))
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
