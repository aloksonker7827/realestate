package impl

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	"gitlab.com/umi7/DezHab_user_backend/service"
	"net/http"
)

const (
	DataBaseStatus = "DataBaseStatus"
	// green means healthy status
	green = "green"
	// red means unhealthy status
	red = "red"
)

// NewHealthCheckResponse sets the default values for dto.HealthCheckResponse
func NewHealthCheckResponse() (int, dto.HealthCheckResponse) {
	healthCheckResponse := dto.HealthCheckResponse{
		StatusCode: red,
		Data: map[string]string{
			DataBaseStatus: red,
		},
	}
	code := http.StatusInternalServerError
	return code, healthCheckResponse
}

var HealthCheck service.HealthCheck

type HealthCheckImpl struct{}

func init() {
	HealthCheck = HealthCheckImpl{}
}

// PostgresDB checks health status of postGreSQL database
func (h HealthCheckImpl) PostgresDB(ctx context.Context) (err error) {
	// calling ping function for checking database health status
	err = impl.Write.Ping()
	return
}

// DeepHealthCheck checks the health status of dependencies of dezhab user
func (h HealthCheckImpl) DeepHealthCheck(ctx context.Context) (code int, healthCheckResponse dto.HealthCheckResponse) {
	// calling PostgresDB to check status of DB
	dbError := h.PostgresDB(ctx)
	// setting default values to code and healthCheckResponse
	code, healthCheckResponse = NewHealthCheckResponse()
	if dbError == nil {
		healthCheckResponse.StatusCode = green
		healthCheckResponse.Data[DataBaseStatus] = green
		code = http.StatusOK
		return
	}
	// if program reaches here then there is some failure in dependencies
	code = http.StatusServiceUnavailable
	return
}
