package routes

import (
	httpDelivery "go-clean-api/internal/delivery/http/handler"
	"go-clean-api/internal/initialize"
	"go-clean-api/internal/repository"
	"go-clean-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, container *initialize.Container) {

	api := e.Group("/api")
	v1 := api.Group("/v1")
	conf := container.Config
	logger := container.ZapLogger

	authHandler := httpDelivery.NewAuthHandler(conf)

	userRepo := repository.NewUserRepository(container.DbGolangSimpleApi)
	userUsecase := usecase.NewUserUsecase(userRepo, logger)
	userHandler := httpDelivery.NewUserHandler(userUsecase, container)

	RegisterAuthRoutes(v1, authHandler)
	RegisterUserRoutes(v1, userHandler)
}
