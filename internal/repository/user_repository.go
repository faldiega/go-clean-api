package repository

import (
	"context"
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/repository"
	"go-simple-api/pkg/common/pagination"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll(ctx context.Context, p pagination.Pagination) ([]entity.User, int, error) {
	var users []entity.User
	var totalItems int64

	// hitung total dulu
	if err := r.db.Model(&entity.User{}).Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	// query dengan limit & offset
	if err := r.db.
		Limit(p.Limit).
		Offset(p.Offset()).
		Where("is_active = ?", true).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(totalItems), nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User

	err := r.db.First(&user, id).Error

	return user, err
}

func (r *userRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	user = entity.User{
		Name:        user.Name,
		Email:       user.Email,
		IsActive:    user.IsActive,
		CreatedDate: time.Now(),
	}
	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) Update(ctx context.Context, user entity.User) (entity.User, error) {
	updatedDate := time.Now()
	user.UpdatedDate = &updatedDate

	err := r.db.
		Model(&entity.User{}).
		Where("id = ?", user.ID).
		Select("Name", "Email", "IsActive", "UpdatedDate").
		Updates(user).Error

	return user, err
}

func (r *userRepository) Delete(ctx context.Context, id int) error {

	return r.db.Delete(&entity.User{}, id).Error
}
