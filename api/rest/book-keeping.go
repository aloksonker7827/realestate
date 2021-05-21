package rest

import (
	"context"
	"encoding/json"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/api"
	"gitlab.com/umi7/DezHab_user_backend/api/handler"
	"gitlab.com/umi7/DezHab_user_backend/common"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	errorHandler "gitlab.com/umi7/DezHab_user_backend/error-handler"
	"gitlab.com/umi7/DezHab_user_backend/service/impl"
	"net/http"
)

const (
	TypeErr = "server error occurred"
)

var GetBookKeeping = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	getBookKeepingRequest, ok := request.(*dto.GetBookKeepingRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	bookKeeping, err := impl.BookKeeping.Get(ctx, getBookKeepingRequest.RequestID)
	if err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	// return the response
	return handler.Response{
		Code: http.StatusOK,
		Payload: dto.BookKeepingResponse{
			Response: dto.BookKeepingStatus{
				Status: bookKeeping.Status,
				Reason: bookKeeping.FailureReason,
				Action: bookKeeping.FailureAction,
			},
		},
	}
}

var UpdateBookKeeping = func(ctx context.Context, req *http.Request, request common.Validator) handler.Response {
	updateBookKeepingRequest, ok := request.(*dto.UpdateBookKeepingRequest)
	if !ok {
		return api.HandleError(ctx, errors.New(TypeErr), errorHandler.ServerErr)
	}

	data, _ := json.Marshal(updateBookKeepingRequest.Data)
	if err := impl.BookKeeping.Update(ctx,
		updateBookKeepingRequest.RequestID, updateBookKeepingRequest.Status,
		updateBookKeepingRequest.Message, string(data)); err != nil {
		return api.HandleMerchantError(ctx, err)
	}

	return handler.Response{
		Code: http.StatusOK,
	}
}
