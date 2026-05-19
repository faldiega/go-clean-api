package middleware

import (
	"fmt"
	"go-simple-api/pkg/common/response"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type RecoveryMiddleware struct {
	logger *zap.Logger
}

func NewRecoveryMiddleware(log *zap.Logger) *RecoveryMiddleware {
	return &RecoveryMiddleware{logger: log}
}

func (m *RecoveryMiddleware) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			defer func() {
				if r := recover(); r != nil {

					// log detail panic ke zap
					m.logger.Error("PANIC RECOVERED",
						zap.String("panic", fmt.Sprintf("%v", r)),
						zap.String("stack", string(debug.Stack())),
						zap.String("method", c.Request().Method),
						zap.String("path", c.Request().URL.Path),
					)

					// return response GSA-500 ke client
					response.SendError(c, http.StatusInternalServerError, "internal server error", "something went wrong, please try again later")
				}
			}()

			return next(c)
		}
	}
}
