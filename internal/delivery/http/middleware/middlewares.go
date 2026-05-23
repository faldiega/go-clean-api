package middleware

import (
	"go-simple-api/internal/config"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Middlewares struct {
	Auth     *JWTMiddleware
	Recovery *RecoveryMiddleware
	Trace    *TraceMiddleware
	Timeout  *TimeoutMiddleware
}

func Load(conf *config.Config, log *zap.Logger) *Middlewares {

	// Auth Middleware
	authMiddleware := NewJWTMiddleware(
		conf.Jwt.Secret,
		conf.WhitelistURLs,
	)

	// Recovery Middleware
	recoveryMiddleware := NewRecoveryMiddleware(log)

	// Trace Middleware
	traceMiddleware := NewTraceMiddleware(log)

	// Timeout Middleware
	timeout := time.Duration(conf.App.Timeout) * time.Second
	timeoutMiddleware := NewTimeoutMiddleware(timeout, log)

	return &Middlewares{
		Auth:     authMiddleware,
		Recovery: recoveryMiddleware,
		Trace:    traceMiddleware,
		Timeout:  timeoutMiddleware,
	}
}

func RegisterMiddlewares(e *echo.Echo, m *Middlewares) {
	e.Use(m.Recovery.Middleware())
	e.Use(m.Trace.Middleware())
	e.Use(m.Timeout.Middleware())
	e.Use(m.Auth.Middleware())
}
