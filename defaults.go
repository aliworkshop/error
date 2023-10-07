package error

var DefaultUnAuthorizedError = UnAuthorized()

var DefaultForbiddenError = Forbidden()

var DefaultValidationError = Validation()

var DefaultNotFoundError = NotFound()

var DefaultInternalError = Internal()

var DefaultDuplicateError = Validation().
	WithId("DuplicateEntryError").
	WithMessage("Requested information already exists").
	SetDefaults(true)
