package usecase

import (
	"context"
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/repository"
	"go-simple-api/internal/domain/usecase"
	pkgLogger "go-simple-api/pkg/common/logger"
	"go-simple-api/pkg/common/pagination"

	"go.uber.org/zap"
	"gorm.io/gorm/utils"
)

type userUsecase struct {
	repo   repository.UserRepository
	logger *zap.Logger
}

func NewUserUsecase(r repository.UserRepository, log *zap.Logger) usecase.UserUsecase {
	return &userUsecase{repo: r, logger: log}
}

func (u *userUsecase) GetUsers(ctx context.Context, p pagination.Pagination) ([]entity.User, int, error) {
	logger := pkgLogger.FromContext(ctx, u.logger)

	logger.Info("fetching all users")

	users, total, err := u.repo.FindAll(ctx, p)
	if err != nil {
		logger.Error("failed to fetch users from repository", zap.Error(err))
		return nil, 0, err
	}
	return users, total, nil
}

func (u *userUsecase) GetUser(ctx context.Context, id int) (entity.User, error) {
	logger := pkgLogger.FromContext(ctx, u.logger)
	var ent entity.User = entity.User{}

	logger.Info("fetching user by id")

	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		logger.Error("failed to get user by id from repository", zap.Error(err))
		return ent, err
	}

	return user, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	logger := pkgLogger.FromContext(ctx, u.logger)

	logger.Info("create new user")

	user, err := u.repo.Create(ctx, user)
	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		return user, err
	}

	return user, err
}

func (u *userUsecase) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	logger := pkgLogger.FromContext(ctx, u.logger)

	logger.Info("update user")

	user, err := u.repo.Update(ctx, user)
	if err != nil {
		logger.Error("failed to update user", zap.Error(err))
		return user, err
	}

	return user, err
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int) error {
	logger := pkgLogger.FromContext(ctx, u.logger)

	logger.Info("delete user with id: " + utils.ToString(id))

	err := u.repo.Delete(ctx, id)
	if err != nil {
		logger.Error("failed to delete user", zap.Error(err))
		return err
	}

	return nil
}
