package repository

import (
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/repository"
	"time"

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

	err := r.db.Where("is_active = ?", true).Find(&users).Error

	return users, err
}

func (r *userRepository) FindByID(id int) (entity.User, error) {
	var user entity.User

	err := r.db.First(&user, id).Error

	return user, err
}

func (r *userRepository) Create(user entity.User) (entity.User, error) {
	user = entity.User{
		Name:        user.Name,
		Email:       user.Email,
		IsActive:    user.IsActive,
		CreatedDate: time.Now(),
	}
	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) Update(user entity.User) (entity.User, error) {
	updatedDate := time.Now()
	user.UpdatedDate = &updatedDate

	err := r.db.
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Select("Name", "Email", "IsActive", "UpdatedDate").
		Updates(user).Error

	return user, err
}

func (r *userRepository) Delete(id int) error {

	return r.db.Delete(&entity.User{}, id).Error
}
