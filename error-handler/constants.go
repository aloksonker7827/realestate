package error_handler

import "net/http"

const (
	AuthenticationError     = "No authentication provided"
	AuthorizationError      = "No authorization provided"
	ValidationError         = "Input validation failed"
	InternalServerError     = "Server error occurred"
	BadRequestError         = "Bad request"
	FailedDependencyError   = "Dependency failed"
	PayloadTooLargeError    = "Payload too large"
	StatusExpectationFailed = "StatusExpectationFailed"
	NoContentError          = "No Mapping Found"
)

const (
	MerchantNotFound = "merchant not found in user-service"
	PartialFailure   = "Partial failure because some of the ids are not found in the db"
	TotalFailure     = "Total failure because none of the ids are found in the db"
)

var ErrorCodeToHttpCode = map[string]int{
	AuthenticationError:     http.StatusForbidden,
	AuthorizationError:      http.StatusUnauthorized,
	ValidationError:         http.StatusUnprocessableEntity,
	InternalServerError:     http.StatusInternalServerError,
	BadRequestError:         http.StatusBadRequest,
	FailedDependencyError:   http.StatusFailedDependency,
	PayloadTooLargeError:    http.StatusRequestEntityTooLarge,
	StatusExpectationFailed: http.StatusExpectationFailed,
	NoContentError:          http.StatusNoContent,
}
