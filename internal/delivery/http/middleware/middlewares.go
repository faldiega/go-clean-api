package middleware

import (
	"go-simple-api/internal/config"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Middlewares struct {
	Auth     *JWTMiddleware
	Recovery *RecoveryMiddleware
	Trace    *TraceMiddleware
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

	return &Middlewares{
		Auth:     authMiddleware,
		Recovery: recoveryMiddleware,
		Trace:    traceMiddleware,
	}
}

func RegisterMiddlewares(e *echo.Echo, m *Middlewares) {
	e.Use(m.Recovery.Middleware())
	e.Use(m.Trace.Middleware())
	e.Use(m.Auth.Middleware())
}
