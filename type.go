package error

const (
	TypeValidation       ErrorType = "VALIDATION"
	TypeNotFound                   = "NOT_FOUND"
	TypeUnAuthorized               = "UNAUTHORIZED"
	TypeForbidden                  = "FORBIDDEN"
	TypeInternal                   = "INTERNAL"
	TypeTooManyRequests            = "TOO_MANY_REQUESTS"
	TypeFailedDependency           = "FAILED_DEPENDENCY"
	TypeTooEarly                   = "TOO_EARLY"
)

type ErrorType string

func (error *err) Is(errType ErrorType) bool {
	return error.Type() == errType
}

func Is(err error, errType ErrorType) bool {
	if e, ok := err.(ErrorModel); ok {
		return e.Is(errType)
	}
	return false
}
