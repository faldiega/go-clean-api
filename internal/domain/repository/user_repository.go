package repository

import (
	"context"
	"go-simple-api/internal/domain/entity"
	domain "go-simple-api/internal/domain/entity"
	"go-simple-api/pkg/common/pagination"
)

type UserRepository interface {
	FindAll(ctx context.Context, p pagination.Pagination) ([]domain.User, int, error)
	FindByID(ctx context.Context, id int) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, id int, updates map[string]interface{}) (entity.User, error)
	Delete(ctx context.Context, id int) error
}
