package main

import (
	"go-simple-api/config"
	httpDelivery "go-simple-api/internal/delivery/http"
	"go-simple-api/internal/repository"
	"go-simple-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	logger, err := config.InitLogger()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	logger.Info("application starting")

	db := config.InitDB()
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	httpDelivery.NewUserHandler(e, userUsecase)

	logger.Info("server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
