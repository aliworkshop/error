package error

import "reflect"

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

	isComparable := reflect.TypeOf(errType).Comparable()
	if isComparable && err == errType {
		return true
	}

	return false
}
