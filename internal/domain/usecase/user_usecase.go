package usecase

import (
	"context"
	"go-simple-api/internal/domain/entity"
	"go-simple-api/pkg/common/pagination"
)

type UserUsecase interface {
	GetUsers(ctx context.Context, p pagination.Pagination) ([]entity.User, int, error)
	GetUser(ctx context.Context, id int) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, id int, updates map[string]interface{}) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}
