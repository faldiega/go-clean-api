package routes

import (
	httpDelivery "go-simple-api/internal/delivery/http/handler"
	"go-simple-api/internal/initialize"
	"go-simple-api/internal/repository"
	"go-simple-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(v1 *echo.Group, container *initialize.Container) {

	httpDelivery.NewAuthHandler(v1, container.Config)

	userRepo := repository.NewUserRepository(container.DbGolangSimpleApi)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := httpDelivery.NewUserHandler(userUsecase)

	RegisterUserRoutes(v1, userHandler)
}
