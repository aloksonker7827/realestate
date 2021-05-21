package error_handler

// Error is the default error message that is generated.
// This is used for all the error cases.
//
// swagger:response error
type Error struct {
	Code int `json:"code"`
	// Message denotes the error message in high-level
	//
	// in: body
	// Required: true
	// Example: Input Validation Failed
	Message string `json:"message"`
	// Data denotes the actual error that is happened in the system
	//
	// in: body
	// Required: true
	// Example: Nil pointer exception
	Data string `json:"data"`
}

func (e Error) Error() string {
	return e.Message
}
