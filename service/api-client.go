package service

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dto"
)

type ApiClient interface {
	GetClient(ctx context.Context, apiKey string) models.ApiClient
	GetAllClients(ctx context.Context) []models.ApiClient
	CreateClient(ctx context.Context, client dto.ApiClient) (dto.ApiClient, error)
	GetSubTagClient(ctx context.Context, apiClient string) string
}

//go:generate mockgen -source=api-client.go -destination=mock-impl/api-client-impl.go -package=mock_impl
