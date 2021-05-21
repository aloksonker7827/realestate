package constants

const (
	RequestData           = "request-data"
	RequestID             = "request-id"
	GoPath                = "GOPATH"
	Provider              = "provider"
	UploadedBy            = "uploadedBy"
	Action                = "action"
	SubTag                = "subTag"
	ClientID              = "clientId"
	Result                = "result"
	Environment           = "ENV"
	DefaultEnvironment    = "dev"
	ProductionEnvironment = "prod"
)

// error messages
const (
	InvalidOrProcessedRequestID = "invalid request id"
	MaxStoreLimitExceededErr    = "max limit exceeded - Max %d stores are allowed at once"
	MaxStoreRecordExceededErr   = "max limit exceeded - Max %d stores and %d records are allowed at once"
	UnauthorizedAPIKey          = "unauthorized API Key"
	InvalidAction               = "invalid action provided in request"
	FilterInvalidFormat         = "applied filters are not in valid format"
	ActionRequiredErr           = "action filed is required"
)
