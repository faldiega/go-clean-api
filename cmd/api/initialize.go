package main

import (
	"go-simple-api/internal/config"
	"go-simple-api/internal/delivery/http"
	"go-simple-api/internal/infrastructure/database"
	"go-simple-api/internal/repository"
	"go-simple-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func Initialize(e *echo.Echo, cfg *config.Config) error {

	db, err := database.NewDatabaseConfig(cfg)
	if err != nil {
		return err
	}

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	http.NewUserHandler(e, userUsecase)

	return nil
}
