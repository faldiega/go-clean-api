package main

import (
	"go-simple-api/internal/config"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	conf := config.LoadConfig()

	use := Initialize(conf)
	use.InitLogger()
	use.Database()
	use.DependencyInjection(e)

	log := use.ZapLogger

	defer log.Sync()
	log.Info("application starting")
	log.Info("server started on port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
