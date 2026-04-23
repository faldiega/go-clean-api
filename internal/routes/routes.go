package routes

import (
	httpDelivery "go-simple-api/internal/delivery/http"
	"go-simple-api/internal/initialize"
	"go-simple-api/internal/repository"
	"go-simple-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(v1 *echo.Group, container *initialize.Container) {

	userRepo := repository.NewUserRepository(container.DbGolangSimpleApi)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := httpDelivery.NewUserHandler(userUsecase)

	httpDelivery.RegisterUserRoutes(v1, userHandler)
}
