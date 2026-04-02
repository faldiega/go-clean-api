package main

import (
	"go-simple-api/internal/config"
	"go-simple-api/internal/infrastructure/logger"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	conf := config.LoadConfig()

	log, err := logger.InitLogger(conf)
	if err != nil {
		panic(err)
	}

	defer log.Sync()

	log.Info("application starting")

	err = Initialize(e, conf)
	if err != nil {
		panic(err)
	}

	log.Info("server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))

}
