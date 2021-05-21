package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

// Provider contains the details of the provider being added like
// client name, provider name, subtag.
//
// swagger:model
type ApiClient struct {
	// StoreID denotes the ID of the store for which
	// the menu needs to be ingested.
	//
	// Provider contains the name of the provider
	// required: true
	Provider string `json:"provider" validate:"required"`
	// Client contains the name of the client
	// required: true
	Client string `json:"client" validate:"required"`
	// Subtag contains the name of the Provider Subtag
	SubTag  string `json:"subtag"`
	ApiKey  string `json:"apikey"`
	Enabled bool   `json:"enabled"`
}

// ApiClientResponse
type ApiClientResponse struct {
	// Data will return an array of api_clients
	Code string             `json:"code"`
	Data []models.ApiClient `json:"data"`
}
