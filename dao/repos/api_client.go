package repos

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

type ApiClient interface {
	Create(context.Context, models.ApiClient) (models.ApiClient, error)
	Search(context.Context, string) (models.ApiClient, error)
	FindAll(context.Context) ([]models.ApiClient, error)
	SearchSubTag(context.Context, string) (string, error)
}

// Command to generate the mock implementation of the ApiClient interface
//go:generate mockgen -destination=mock-impl/api_client_impl.go -package=mock_impl -source=api_client.go ApiClient
