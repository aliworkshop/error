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
