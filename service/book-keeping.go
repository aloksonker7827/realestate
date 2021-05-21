package service

import (
	"context"

	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

type BookKeeping interface {
	Create(ctx context.Context, status string) error
	Update(ctx context.Context, requestId, status, failureReason, failureAction string) error
	Get(ctx context.Context, requestId string) (models.BookKeeping, error)
	CheckBookKeepingRecord(ctx context.Context, requestId string) (models.BookKeeping, error)
	FetchPaginated(ctx context.Context, action string, page int, pageSize int, order string) (bookKeeping []models.BookKeeping, err error)
	Count(ctx context.Context, action string) (count int, err error)
}

// Command to generate the mock implementation of the BookKeeping interface
//go:generate mockgen -source=book-keeping.go -destination=mock-impl/book_keeping_impl.go -package=mock_impl
