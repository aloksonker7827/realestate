package common

// Asynchronous is used to determine the requests which wil be processed asynchronously.
//
// If the request dto is of type Asynchronous, some global-level functionality are
// added in the handler.
type Asynchronous interface {
	// AsyncProcessor can be implemented on the respective dto to do
	// additional task based on the nature of the API.
	// It is used to perform some operation for Asynchronous API.
	//
	// This function can be used to perform any other concurrent functionality
	// for a particular asynchronous API.
	//
	// Note : When a particular dto implements this function and follows the rest standard,
	// book-keeping functionality is added automatically in the rest handler.
	//
	// Book-Keeping functionality denotes that, for a given asynchronous API,
	// the entire request related information is saved in the DB against a requestId. RequestId is a uniqueId which
	// will be generated for each request. Most of the asynchronous API, returns the callback-url
	// using which caller can know about the status of the request. This status will be updated in the
	// DB by book-keeping-service once the request is processed by the appropriate processor
	AsyncProcessor()
}
