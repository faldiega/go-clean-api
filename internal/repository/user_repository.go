package repository

import (
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]entity.User, error) {

	var users []entity.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) FindByID(id uint) (entity.User, error) {

	var user entity.User

	err := r.db.First(&user, id).Error

	return user, err
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {

	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) Update(user entity.User) (entity.User, error) {

	err := r.db.Save(&user).Error

	return user, err
}

func (r *userRepository) Delete(id uint) error {

	return r.db.Delete(&entity.User{}, id).Error
}
