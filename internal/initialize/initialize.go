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
	Auth              *middleware.JWTMiddleware
}

func NewContainer() *Container {

	conf := config.LoadConfig()

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

	// Middleware
	jwtMiddleware := middleware.NewJWTMiddleware(
		conf.Jwt.JwtSecret,
		conf.WhitelistURLs,
	)

	return &Container{
		Config:            conf,
		ZapLogger:         log,
		DbGolangSimpleApi: db,
		Auth:              jwtMiddleware,
	}

}
