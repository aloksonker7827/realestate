package repos

import (
	"context"

	"gitlab.com/umi7/DezHab_user_backend/dao/models"
)

type BookKeeping interface {
	Create(context.Context, models.BookKeeping) error
	Update(context.Context, models.BookKeeping) error
	Search(context.Context, string) (models.BookKeeping, error)
	FetchPaginated(context.Context, string, int, int, string) ([]models.BookKeeping, error)
	Count(context.Context, string) (int, error)
}

// Command to generate the mock implementation of the BookKeeping interface
//go:generate mockgen -source=book_keeping.go -destination=mock-impl/book_keeping_impl.go -package=mock_impl
