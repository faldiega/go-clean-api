package usecase

import "go-simple-api/internal/domain/entity"

type UserUsecase interface {
	GetUsers() ([]entity.User, error)
	GetUser(id int) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id int) error
}
