package impl

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	"time"

	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos/impl"
	"gitlab.com/umi7/DezHab_user_backend/service"
)

var ApiClient service.ApiClient

type ApiClientImpl struct{}

func init() {
	ApiClient = ApiClientImpl{}
}

func (a ApiClientImpl) GetClient(ctx context.Context, apiKey string) (client models.ApiClient) {
	client, err := impl.ApiClient.Search(ctx, apiKey)
	if err != nil {
		return
	}
	return
}

func (a ApiClientImpl) GetAllClients(ctx context.Context) (client []models.ApiClient) {
	client, err := impl.ApiClient.FindAll(ctx)
	if err != nil {
		return
	}
	return
}

func (a ApiClientImpl) convertDtoToModel(ctx context.Context, client dto.ApiClient) (apiClient models.ApiClient) {
	return models.ApiClient{
		Client:     client.Client,
		ApiKey:     client.ApiKey,
		Provider:   client.Provider,
		SubTag:     client.SubTag,
		Enabled:    client.Enabled,
		MaxRetries: 0,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		DeletedAt:  nil,
	}
}

func (a ApiClientImpl) convertModelToDto(ctx context.Context, apiClient models.ApiClient) (client dto.ApiClient) {
	return dto.ApiClient{
		Provider: apiClient.Provider,
		Client:   apiClient.Client,
		SubTag:   apiClient.SubTag,
		ApiKey:   apiClient.ApiKey,
		Enabled:  apiClient.Enabled,
	}
}

func (a ApiClientImpl) CreateClient(ctx context.Context, client dto.ApiClient) (apiClient dto.ApiClient, err error) {
	client.ApiKey = uuid.New().String()
	client.Enabled = true
	modelClient, err := impl.ApiClient.Create(ctx, a.convertDtoToModel(ctx, client))
	if err != nil {
		return
	}
	apiClient = a.convertModelToDto(ctx, modelClient)
	return
}

func (a ApiClientImpl) GetSubTagClient(ctx context.Context, apiClient string) (subTag string) {
	subTag, err := impl.ApiClient.SearchSubTag(ctx, apiClient)
	if err != nil {
		return
	}
	return
}
