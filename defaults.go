package error

var DefaultUnAuthorizedError = UnAuthorized().
	WithId("UnAuthorizedError").
	WithMessage("You are not authorized.")

var DefaultForbiddenError = Forbidden().
	WithId("ForbiddenError").
	WithMessage("Access to this section is denied.")

var DefaultValidationError = Validation().
	WithId("ValidationError").
	WithMessage("Invalid request information.")

var DefaultNotFoundError = NotFound().
	WithId("NotFoundError").
	WithMessage("Requested information not found.")

var DefaultInternalError = Internal().
	WithId("InternalError").
	WithMessage("Internal error occurred.")

var DefaultDuplicateError = Validation().
	WithId("DuplicateEntryError").
	WithMessage("Requested information already exists")
