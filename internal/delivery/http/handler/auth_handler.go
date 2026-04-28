package handler

import (
	"net/http"

	"go-simple-api/internal/config"
	"go-simple-api/internal/delivery/http/dto"
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

	key := h.conf.AppSecret

	var req dto.TokenRequest
	if err := c.Bind(&req); err != nil {
		return response.SendError(c, http.StatusBadRequest, "invalid request", err.Error())
	}

	if req.ApiKey != key {
		return response.SendError(c, http.StatusUnauthorized, "you don't have access to this API", error.Error(echo.ErrUnauthorized))
	}

	data, err := infrastructure.GenerateToken(userID, h.conf)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "failed generate token", err.Error())
	}

	result := dto.TokenResponse{
		Token:     data.Token,
		Type:      "Bearer",
		ExpiredAt: data.ExpiredAt,
	}

	return response.SendSuccess(c, http.StatusOK, "successfully", result)
}
