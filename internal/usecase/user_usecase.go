package usecase

import (
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/repository"
	"go-simple-api/internal/domain/usecase"
	"go-simple-api/pkg/common/pagination"
)

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{r}
}

func (u *userUsecase) GetUsers(p pagination.Pagination) ([]entity.User, int, error) {
	return u.repo.FindAll(p)
}

func (u *userUsecase) GetUser(id int) (entity.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecase) CreateUser(user entity.User) (entity.User, error) {
	return u.repo.Create(user)
}

func (u *userUsecase) UpdateUser(user entity.User) (entity.User, error) {
	return u.repo.Update(user)
}

func (u *userUsecase) DeleteUser(id int) error {
	return u.repo.Delete(id)
}
