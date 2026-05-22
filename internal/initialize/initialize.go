package initialize

import (
	"go-simple-api/internal/config"
	"go-simple-api/internal/delivery/http/middleware"
	"go-simple-api/internal/infrastructure/database"
	"go-simple-api/internal/infrastructure/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	Config            *config.Config
	DbGolangSimpleApi *gorm.DB
	ZapLogger         *zap.Logger
	Middlewares       *middleware.Middlewares
}

func NewContainer() *Container {

	// Load configuration
	conf := config.Load()

	// Logger
	log, err := logger.ZapLogger(conf)
	if err != nil {
		panic(err)
	}

	// Database
	db, err := database.NewDatabaseConfig(conf)
	if err != nil {
		panic(err)
	}

	// Load middleware
	middlewares := middleware.Load(conf, log)

	return &Container{
		Config:            conf,
		ZapLogger:         log,
		DbGolangSimpleApi: db,
		Middlewares:       middlewares,
	}

}
