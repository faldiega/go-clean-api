package server

import (
	"context"
	"errors"
	"fmt"
	"go-simple-api/internal/delivery/http/routes"
	"go-simple-api/internal/initialize"
	"go-simple-api/pkg/common/constants"
	"go-simple-api/pkg/common/graceful"
	"go-simple-api/pkg/common/validator"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func NewEchoApp(container *initialize.Container) *echo.Echo {

	e := echo.New()

	// hide default Echo banner & error handler
	e.HideBanner = true
	e.HidePort = true

	// register custom validator
	e.Validator = validator.GetValidator()

	// register middleware
	e.Use(container.Recovery.Middleware())
	e.Use(container.Auth.Middleware())

	routes.RegisterRoutes(e, container)

	return e
}

func Start(container *initialize.Container) {

	conf := container.Config
	log := container.ZapLogger
	app := NewEchoApp(container)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", conf.App.Port),
		Handler:           app,
		WriteTimeout:      constants.DefaultServerWriteTimeout,
		ReadTimeout:       constants.DefaultServerReadTimeout,
		IdleTimeout:       constants.DefaultServerIdleTimeout,
		ReadHeaderTimeout: constants.DefaultServerReadHeaderTimeout,
	}

	go func() {
		if err := app.StartServer(srv); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Error start server", zap.Error(err))
		}
	}()

	log.Info("application starting")
	log.Info("server started on port 8080")
	log.Sync() // Force flush to ensure logs are written immediately

	wait := graceful.Shutdown(context.Background(), 10*time.Second, map[string]graceful.Operation{
		"http-server": func(ctx context.Context) error {
			return app.Shutdown(ctx)
		},
	})
	<-wait
}
