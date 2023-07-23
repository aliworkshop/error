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

// FIXME remove NotFoundError
var NotFoundError = NotFound().
	WithId("NotFoundError").
	WithMessage("Requested information not found.")

var DefaultNotFoundError = NotFoundError

var DefaultInternalError = Internal().
	WithId("InternalError").
	WithMessage("Internal error occurred.")

var DefaultDuplicateError = Validation().
	WithId("DuplicateEntryError").
	WithMessage("Requested information already exists")

var DefaultTooManyRequestsError = Validation().
	WithId("TooManyRequestsError").
	WithMessage("You've made too many requests recently. Please wait and try your request again later.")

var DefaultFailedDependencyError = FailedDependency().
	WithId("FailedDependencyError").
	WithMessage("The request could not be performed on the resource because the requested action depended on another action and that action failed.")
