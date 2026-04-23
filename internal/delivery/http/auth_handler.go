package http

import (
	"net/http"

	"go-simple-api/internal/config"
	"go-simple-api/internal/infrastructure"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	conf *config.Config
}

func NewAuthHandler(e *echo.Echo, cfg *config.Config) {
	handler := &AuthHandler{cfg}

	e.POST("/auth", handler.Login)
}

func (h *AuthHandler) Login(c echo.Context) error {

	// nanti ini diganti pakai request body / validasi DB
	userID := uint(1)

	token, err := infrastructure.GenerateToken(userID, h.conf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
