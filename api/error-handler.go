package api

import (
	"context"
	"gitlab.com/umi7/DezHab_user_backend/api/handler"
	errorhandler "gitlab.com/umi7/DezHab_user_backend/error-handler"
)

var HandleError = func(ctx context.Context, err error, f func(err error) errorhandler.Error) handler.Response {
	errorResponse := f(err)
	return handler.Response{
		Code:    errorResponse.Code,
		Payload: errorResponse,
	}
}

var HandleMerchantError = func(ctx context.Context, err error) handler.Response {
	if err.Error() == errorhandler.PartialFailure || err.Error() == errorhandler.TotalFailure {
		return HandleError(ctx, err, errorhandler.MissingDataError)
	}
	if err.Error() == errorhandler.MerchantNotFound {
		return HandleError(ctx, err, errorhandler.ExpectationFailedErr)
	} else {
		return HandleError(ctx, err, errorhandler.ServerErr)
	}
}