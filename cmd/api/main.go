package main

import (
	"go-simple-api/config"
	httpDelivery "go-simple-api/internal/delivery/http"
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/repository"
	"go-simple-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	db := config.InitDB()
	db.AutoMigrate(&entity.User{})

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	httpDelivery.NewUserHandler(e, userUsecase)

	e.Logger.Fatal(e.Start(":8080"))
}
