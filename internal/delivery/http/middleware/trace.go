package middleware

import (
	"go-simple-api/pkg/common/constants"
	pkgLogger "go-simple-api/pkg/common/logger"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type TraceMiddleware struct {
	logger *zap.Logger
}

func NewTraceMiddleware(log *zap.Logger) *TraceMiddleware {
	return &TraceMiddleware{logger: log}
}

func (m *TraceMiddleware) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			start := time.Now()

			// 1. generate trace ID
			traceID := uuid.New().String()

			// 2. buat logger dengan trace ID
			traceLogger := pkgLogger.WithTraceID(m.logger, traceID)

			// 3. simpan trace ID ke Echo context (untuk header response)
			c.Set(string(constants.TraceIDKey), traceID)

			// 4. simpan logger ke request context (untuk usecase/repository)
			ctx := pkgLogger.WithLogger(c.Request().Context(), traceLogger)
			c.SetRequest(c.Request().WithContext(ctx))

			// set Trace ID ke response header (opsional)
			c.Response().Header().Set("X-Trace-ID", traceID)

			// 5. log request masuk
			traceLogger.Info("request started",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.String("ip", c.RealIP()),
			)

			// 6. lanjut ke handler
			err := next(c)

			// 7. log request selesai
			traceLogger.Info("request finished",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.Int("status", c.Response().Status),
				zap.Duration("latency", time.Since(start)),
			)

			return err
		}
	}
}
