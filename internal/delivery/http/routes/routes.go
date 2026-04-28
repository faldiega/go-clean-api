package routes

import (
	httpDelivery "go-simple-api/internal/delivery/http/handler"
	"go-simple-api/internal/initialize"
	"go-simple-api/internal/repository"
	"go-simple-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, container *initialize.Container) {

	api := e.Group("/api")
	v1 := api.Group("/v1")

	authHandler := httpDelivery.NewAuthHandler(container.Config)

	userRepo := repository.NewUserRepository(container.DbGolangSimpleApi)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := httpDelivery.NewUserHandler(userUsecase)

	RegisterAuthRoutes(v1, authHandler)
	RegisterUserRoutes(v1, userHandler)
}
