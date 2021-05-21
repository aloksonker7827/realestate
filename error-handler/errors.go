package error_handler

import "encoding/json"

type ErrorHandler func(error) Error

var AuthenticationErr = func(err error) Error {
	return createError(AuthenticationError, err)
}

var AuthorizationErr = func(err error) Error {
	return createError(AuthorizationError, err)
}

var ValidationErr = func(err error) Error {
	return createError(ValidationError, err)
}

var ServerErr = func(err error) Error {
	return createError(InternalServerError, err)
}

var ExpectationFailedErr = func(err error) Error {
	return createError(StatusExpectationFailed, err)
}

var RequestErr = func(err error) Error {
	return createError(BadRequestError, err)
}

var DependencyErr = func(err error) Error {
	return createError(FailedDependencyError, err)
}

var DataErr = func(err error) Error {
	return createError(NoContentError, err)
}

var createError = func(message string, err error) Error {
	data, _ := json.Marshal(err.Error())
	return Error{Code: ErrorCodeToHttpCode[message], Message: message, Data: string(data)}
}

var MissingDataError = func(err error) Error {
	return err.(Error)
}
