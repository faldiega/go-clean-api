package main

import (
	"go-simple-api/internal/config"
	"go-simple-api/internal/delivery/http"
	"go-simple-api/internal/infrastructure/database"
	"go-simple-api/internal/infrastructure/logger"
	"go-simple-api/internal/repository"
	"go-simple-api/internal/usecase"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type globalStruct struct {
	config.Config

	DbGolangSimpleApi *gorm.DB
	ZapLogger         *zap.Logger
}

func Initialize(cfg config.Config) *globalStruct {

	gs := &globalStruct{
		Config: cfg,
	}

	return gs
}

func (gs *globalStruct) InitLogger() {

	log, err := logger.ZapLogger(gs.Config)
	if err != nil {
		panic(err)
	}

	gs.ZapLogger = log
}

func (gs *globalStruct) Database() {

	db, err := database.NewDatabaseConfig(gs.Config)
	if err != nil {
		panic(err)
	}

	gs.DbGolangSimpleApi = db
}

func (gs *globalStruct) EchoMiddleware() {
	/*
		# Take a notes, for using middleware.
		# Still research about this.

		web.InsertFilter("*", web.BeforeRouter, func(ctx *context.Context) {
			middleware.WrapperServer(ctx, ms.telemetryWrapper)
		})

		web.InsertFilterChain("*", middleware.SetDataContext())
		web.InsertFilterChain("*", middleware.SystemAuthentication(ms.platformAuth, ms.platformRedis, ms.platformLog))
		web.InsertFilterChain("*", middleware.CheckUrlUnAuthorized(ms.SkipAuthorized))
		fmt.Println("Success register global middleware")

		web.InsertFilterChain("/api/v1/indomaret/*", indomaretMiddleware.VerifiedHeaderRequest())
		web.InsertFilterChain("/api/v1/indomaret/*", indomaretMiddleware.VerifiedJwtToken(ms.Jwt.Indomaret))
		web.InsertFilterChain("/api/v1/indomaret/*", middleware.SetDataContextPayment(ms.PaymentPointRepo, indomaret.CODE_CHANNEL))
		fmt.Println("Success register indomaret middleware")
	*/
}

func (gs *globalStruct) DependencyInjection(e *echo.Echo) {

	userRepo := repository.NewUserRepository(gs.DbGolangSimpleApi)
	userUsecase := usecase.NewUserUsecase(userRepo)
	http.NewUserHandler(e, userUsecase)
}
