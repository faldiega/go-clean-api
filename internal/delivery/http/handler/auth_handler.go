package handler

import (
	"net/http"

	"go-simple-api/internal/config"
	"go-simple-api/internal/infrastructure"
	"go-simple-api/pkg/common/response"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	conf *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{cfg}
}

func (h *AuthHandler) Login(c echo.Context) error {

	// nanti ini diganti pakai request body / validasi DB
	userID := uint(1)

	token, err := infrastructure.GenerateToken(userID, h.conf)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "failed generate token", err.Error())

	}

	return response.SendSuccess(c, http.StatusOK, "successfully", token)
}
