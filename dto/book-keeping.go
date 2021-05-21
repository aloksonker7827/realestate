package dto

import (
	"gitlab.com/umi7/DezHab_user_backend/common"
	"time"
)

type GetBookKeepingRequest struct {
	RequestID string `json:"requestId" validate:"required"`
}

// Validate function is used to validate the request.
//
// This is the implementation of the Validator interface.
// If any custom validation is required for a dto apart from the
// validation(s) present as part of tags in the struct, this function
// can be used to perform the same.
//
// Note : All the API's which follows rest standard is forced to
// implement this function. As a compliment for all such API's,
// validation will be performed automatically using the common rest handler.
func (g GetBookKeepingRequest) Validate() error {
	return common.Validate.Struct(g)
}

type UpdateBookKeepingRequest struct {
	Code      int         `json:"code"`
	RequestID string      `json:"requestId" validate:"required"`
	Status    string      `json:"status" validate:"required"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

// Validate function is used to validate the request.
//
// This is the implementation of the Validator interface.
// If any custom validation is required for a dto apart from the
// validation(s) present as part of tags in the struct, this function
// can be used to perform the same.
//
// Note : All the API's which follows rest standard is forced to
// implement this function. As a compliment for all such API's,
// validation will be performed automatically using the common rest handler.
func (u UpdateBookKeepingRequest) Validate() error {
	return common.Validate.Struct(u)
}

// Authenticate is used to authenticate the request using api-key
//
// This function can be implemented by the request dto
// to implement the additional function along with the api-key authentication
// based on the nature of the API.
//
// When an API follows the rest standard and the corresponding dto implements the
// Authenticate function, api-key authentication is performed automatically by the rest handler.
// Things to do : This function signature can be changed to return error and perform API specific
// authentication in a appropriate struct.
func (u UpdateBookKeepingRequest) Authenticate() {

}

type BookKeepingResponse struct {
	Response BookKeepingStatus `json:"response"`
}

type BookKeepingStatus struct {
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
	Action string `json:"action,omitempty"`
}

type BookKeepingActionLogs struct {
	RequestID  string    `json:"requestId"`
	Status     string    `json:"status"`
	Reason     string    `json:"reason,omitempty"`
	File       string    `json:"file,omitempty"`
	Action     string    `json:"action,omitempty"`
	UploadedBy string    `json:"uploadedBy,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
