package common

// Authenticator is used to determine the requests needs to be authenticated.
//
// When a dto is of Authenticator type, some global level validations are performed in the request handler
// This interface can be enhanced to support the overall auth for an API by adding Authorization function as well
type Authenticator interface {
	// Authenticate is used to authenticate the request using api-key
	//
	// This function can be implemented by the request dto
	// to implement the additional function along with the api-key authentication
	// based on the nature of the API.
	//
	// When an API follows the rest standard and the corresponding dto implements the
	// Authenticate function, api-key authentication is performed automatically by the rest handler.
	// Things to do : This function signature can be changed to return error and perform API specific
	// authentication in a appropriate struct.
	//
	// Note : When a particular struct implements this interface, the request should have
	// a header with a key as x-api-key and value as the client's API key.
	// This API Key is also used in application to determine the provider in case of
	// food/non-food catalog management API's.
	Authenticate()
}
