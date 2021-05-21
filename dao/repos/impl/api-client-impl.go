package impl

import (
	"context"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos"
)

var ApiClient repos.ApiClient

type ApiClientImpl struct{}

const (
	ErrNoData = "no data found"
)

func init() {
	ApiClient = ApiClientImpl{}
}

func (apiClientImpl ApiClientImpl) Create(ctx context.Context, apiClient models.ApiClient) (client models.ApiClient, err error) {
	if err = impl.Write.Create(&apiClient); err == nil {
		client = apiClient
	}
	return
}

func (apiClientImpl ApiClientImpl) Search(ctx context.Context, apiKey string) (apiClient models.ApiClient, err error) {
	var apiClients []models.ApiClient
	if err = impl.Read.First(&apiClients, models.ApiClient{ApiKey: apiKey, Enabled: true}); err != nil {
		return
	}
	if len(apiClients) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	apiClient = apiClients[0]
	return
}

func (apiClientImpl ApiClientImpl) FindAll(ctx context.Context) (apiClient []models.ApiClient, err error) {
	var apiClients []models.ApiClient
	if err = impl.Read.Find(&apiClients, models.ApiClient{Enabled: true}, models.ApiClient{SubTag: "all"}); err != nil {
		return
	}
	if len(apiClients) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	apiClient = apiClients
	return
}

func (apiClientImpl ApiClientImpl) SearchSubTag(ctx context.Context, provider string) (subTag string, err error) {
	var apiClients []models.ApiClient
	if err = impl.Read.First(&apiClients, models.ApiClient{Provider: provider, Enabled: true}); err != nil {
		return
	}
	if len(apiClients) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	return apiClients[0].SubTag, nil
}
