package dto

// HealthCheckResponse response for deep health check
type HealthCheckResponse struct {
	// StatusCode is the flag sent in response about health status
	// if any dependencies is down StatusCode will be red

	// red means unhealthy
	// green means healthy
	//
	//example: red
	StatusCode string `json:"StatusCode"`
	// Error is not null when service is in error
	Error error `json:"Error"`
	// Data denotes the status of dependencies
	//
	// example: database : red
	Data map[string]string `json:"Data"`
}
