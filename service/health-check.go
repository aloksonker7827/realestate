package service

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dto"
)

// HealthCheck interface checks health status of dezhab user service
// and all its dependencies except external services
//
// This interface's sole purpose is to check whether dezhab user service is running
// properly by checking health status of all its dependencies
type HealthCheck interface {
	DeepHealthCheck(ctx context.Context) (int, dto.HealthCheckResponse)
}
