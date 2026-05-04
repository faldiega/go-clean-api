package repository

import (
	domain "go-simple-api/internal/domain/entity"
	"go-simple-api/pkg/common/pagination"
)

type UserRepository interface {
	FindAll(p pagination.Pagination) ([]domain.User, int, error)
	FindByID(id int) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id int) error
}
