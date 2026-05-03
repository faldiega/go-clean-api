package usecase

import (
	"go-simple-api/internal/domain/entity"
	"go-simple-api/pkg/common/pagination"
)

type UserUsecase interface {
	GetUsers(p pagination.Pagination) ([]entity.User, int, error)
	GetUser(id int) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id int) error
}
