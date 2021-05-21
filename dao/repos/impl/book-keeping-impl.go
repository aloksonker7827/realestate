package impl

import (
	"context"
	"errors"
	"gitlab.com/umi7/DezHab_user_backend/dao/core/impl"
	"gitlab.com/umi7/DezHab_user_backend/dao/models"
	"gitlab.com/umi7/DezHab_user_backend/dao/repos"
	"time"
)

var BookKeeping repos.BookKeeping

type BookKeepingImpl struct{}

const (
	updateWhereQuery = "request_id=?"
)

func init() {
	BookKeeping = BookKeepingImpl{}
}

func (bookKeepingImpl BookKeepingImpl) Create(ctx context.Context, bookKeeping models.BookKeeping) (err error) {
	bookKeeping.CreatedAt = time.Now().UTC()
	bookKeeping.UpdatedAt = bookKeeping.CreatedAt
	err = impl.Write.Create(&bookKeeping)
	return
}

func (bookKeepingImpl BookKeepingImpl) Update(ctx context.Context, bookKeeping models.BookKeeping) (err error) {
	bookKeeping.UpdatedAt = time.Now().UTC()
	updateValues := map[string]interface{}{
		"status":         bookKeeping.Status,
		"failure_reason": bookKeeping.FailureReason,
		"failure_action": bookKeeping.FailureAction,
		"is_retryable":   bookKeeping.IsRetryable,
	}
	err = impl.Write.Update(&bookKeeping, updateWhereQuery, bookKeeping.RequestID, updateValues)
	return
}

func (bookKeepingImpl BookKeepingImpl) Search(ctx context.Context, requestId string) (bookKeeping models.BookKeeping, err error) {
	var bookKeepingList []models.BookKeeping
	if err = impl.Read.Last(&bookKeepingList, models.BookKeeping{RequestID: requestId}); err != nil {
		return
	}
	if len(bookKeepingList) == 0 {
		err = errors.New(ErrNoData)
		return
	}
	bookKeeping = bookKeepingList[0]
	return
}

func (bookKeepingImpl BookKeepingImpl) FetchPaginated(ctx context.Context, action string, pageNum int, pageSize int, order string) (bookKeepingList []models.BookKeeping, err error) {
	limit := pageSize
	offset := pageSize * (pageNum - 1)
	if err = impl.Read.FetchPaginated(&bookKeepingList, map[string]interface{}{"action": action}, models.BookKeeping{}, limit, offset, order); err != nil {
		return
	}
	return
}

func (bookKeepingImpl BookKeepingImpl) Count(ctx context.Context, action string) (count int, err error) {
	var bookKeepingList []models.BookKeeping
	if err = impl.Read.Count(&bookKeepingList, &count, models.BookKeeping{Action: action}, models.BookKeeping{}); err != nil {
		return
	}
	return
}
