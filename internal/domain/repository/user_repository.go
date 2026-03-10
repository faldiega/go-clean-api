package repository

import domain "go-simple-api/internal/domain/entity"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindByID(id uint) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint) error
}
