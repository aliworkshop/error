package error

var DefaultUnAuthorizedError = UnAuthorized()

var DefaultForbiddenError = Forbidden()

var DefaultValidationError = Validation()

var DefaultNotFoundError = NotFound()

var DefaultInternalError = Internal()

var DefaultDuplicateError = Duplicate().
	WithId("DuplicateEntryError").
	WithMessage("Requested information already exists").
	SetDefaults(true)

var DefaultTooManyRequestError = TooManyRequest().
	WithId("TooManyRequestError").
	WithMessage("Maximum requests per seconds has been reached.").
	SetDefaults(true)

var DefaultUnAuthenticatedError = UnAuthorized().
	WithId("MustAuthenticate").
	WithMessage("You must authenticate")
