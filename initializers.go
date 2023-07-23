package error

func Validation(err ...error) ErrorModel {
	var e = New(err...).WithType(TypeValidation)
	return e
}

func Forbidden(err ...error) ErrorModel {
	var e = New(err...).WithType(TypeForbidden)
	return e
}

func Internal(err ...error) ErrorModel {
	var e = New(err...).WithType(TypeInternal)
	return e
}

func NotFound(err ...error) ErrorModel {
	var e = New(err...).WithType(TypeNotFound)
	return e
}

func UnAuthorized(err ...error) ErrorModel {
	var e = New(err...).WithType(TypeUnAuthorized)
	return e
}

func TooManyRequests(err ...error) ErrorModel {
	var e = New(err...).WithType(TypeTooManyRequests)
	return e
}

func FailedDependency(err ...error) ErrorModel {
	var e = New(err...).WithType(TypeFailedDependency)
	return e
}

func TooEarly(err ...error) ErrorModel {
	var e = New(err...).WithType(TypeTooEarly)
	return e
}
