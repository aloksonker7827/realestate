package impl

import (
	"context"
	"errors"
	"fmt"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos/impl"
	"gitlab.com/umi7/DezHab_user_backend/service"
	"gitlab.com/umi7/DezHab_user_backend/utils"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
)

var BookKeeping service.BookKeeping

type BookKeepingImpl struct{}

func init() {
	BookKeeping = BookKeepingImpl{}
}

// Right now, not getting
func (b BookKeepingImpl) Create(ctx context.Context, status string) (err error) {
	bookKeeping := models.BookKeeping{
		ClientID:    utils.GetUIntFromContext(ctx, constants.ClientID),
		Version:     0,
		RequestID:   utils.GetRequestID(ctx),
		RequestData: utils.GetRequestData(ctx),
		Status:      status,
		UploadedBy:  utils.GetUploadedBy(ctx),
		Action:      utils.GetAction(ctx),
		Result:      utils.GetResult(ctx),
	}
	err = impl.BookKeeping.Create(ctx, bookKeeping)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error occurred while creating entry on book-keeping : %v", err))
	}
	return
}

func (b BookKeepingImpl) Update(ctx context.Context, requestId string, status string, failureReason string, failureAction string) (err error) {
	bookKeeping := models.BookKeeping{
		RequestID:     requestId,
		Status:        status,
		FailureReason: failureReason,
		FailureAction: failureAction,
		IsRetryable:   failureReason == "server-error",
	}
	err = impl.BookKeeping.Update(ctx, bookKeeping)
	if err != nil {
		logger.Error(ctx, err)
	}
	return
}

func (b BookKeepingImpl) Get(ctx context.Context, requestId string) (bookKeeping models.BookKeeping, err error) {
	bookKeeping, err = impl.BookKeeping.Search(ctx, requestId)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return
}

func (b BookKeepingImpl) CheckBookKeepingRecord(ctx context.Context, requestId string) (models.BookKeeping, error) {
	// check if request id is valid or not, doesn't need to check for status
	bookKeepingRecord, err := BookKeeping.Get(ctx, requestId)
	if err != nil {
		logger.Error(ctx, "Book keeping record not found: "+err.Error())
		return bookKeepingRecord, errors.New(constants.InvalidOrProcessedRequestID)
	}
	return bookKeepingRecord, nil
}

func (b BookKeepingImpl) FetchPaginated(ctx context.Context, action string, pageNum int, pageSize int, order string) (bookKeeping []models.BookKeeping, err error) {
	bookKeeping, err = impl.BookKeeping.FetchPaginated(ctx, action, pageNum, pageSize, order)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return
}

func (b BookKeepingImpl) Count(ctx context.Context, action string) (count int, err error) {
	count, err = impl.BookKeeping.Count(ctx, action)
	if err != nil {
		logger.Error(ctx, err)
		return
	}
	return
}
