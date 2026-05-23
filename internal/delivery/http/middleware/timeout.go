package middleware

import (
	"context"
	"go-simple-api/pkg/common/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type TimeoutMiddleware struct {
	timeout time.Duration
	logger  *zap.Logger
}

func NewTimeoutMiddleware(timeout time.Duration, log *zap.Logger) *TimeoutMiddleware {
	return &TimeoutMiddleware{
		timeout: timeout,
		logger:  log,
	}
}

func (m *TimeoutMiddleware) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// 1. buat context dengan timeout
			ctx, cancel := context.WithTimeout(
				c.Request().Context(),
				m.timeout,
			)
			defer cancel() // ← wajib! cegah context leak

			// 2. inject context ke request
			c.SetRequest(c.Request().WithContext(ctx))

			// 3. jalankan handler di goroutine terpisah
			done := make(chan error, 1)

			go func() {
				done <- next(c)
			}()

			// 4. tunggu: selesai duluan atau timeout?
			select {
			case err := <-done:
				// handler selesai sebelum timeout
				return err

			case <-ctx.Done():
				// timeout tercapai
				m.logger.Warn("request timeout",
					zap.String("method", c.Request().Method),
					zap.String("path", c.Request().URL.Path),
					zap.Duration("timeout", m.timeout),
				)

				return response.SendError(c, http.StatusRequestTimeout, "request timeout", "request took too long, please try again")
			}
		}
	}
}
