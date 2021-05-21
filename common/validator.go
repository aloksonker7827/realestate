package common

import "gopkg.in/go-playground/validator.v9"

// Validator is used to validate each request's data.
//
// Every API exposed should have a validation on the data being received
type Validator interface {
	// Validate function is used to validate the request.
	//
	// This is the implementation of the Validator interface.
	// If any custom validation is required for a dto apart from the
	// validation(s) present as part of tags in the struct, this function
	// can be used to perform the same.
	//
	// Note : All the API's which follows rest standard is forced to
	// implement this function. As a compliment for all such API's,
	// validation will be performed automatically using the common rest handler.
	Validate() error
}

// Singleton instance to perform the validation
var Validate *validator.Validate

func init() {
	Validate = validator.New()
}
