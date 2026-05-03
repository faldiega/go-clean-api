package handler

import (
	"net/http"

	"go-simple-api/internal/config"
	"go-simple-api/internal/delivery/http/dto"
	"go-simple-api/internal/infrastructure"
	"go-simple-api/pkg/common/response"
	customValidator "go-simple-api/pkg/common/validator"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	apiKey    string
	jwtSecret string
	jwtExpire int
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		apiKey:    cfg.App.Secret,
		jwtSecret: cfg.Jwt.Secret,
		jwtExpire: cfg.Jwt.Expired,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {

	req := new(dto.TokenRequest)
	if err := c.Bind(&req); err != nil {
		return response.SendError(c, http.StatusBadRequest, "invalid request", err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.SendError(c, http.StatusBadRequest, "validation failed", customValidator.FormatValidationError(err))
	}

	if req.ApiKey != h.apiKey {
		return response.SendError(c, http.StatusUnauthorized, "unauthorized", "you don't have access to this API")
	}

	data, err := infrastructure.GenerateToken(h.jwtSecret, h.jwtExpire)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "failed generate token", err.Error())
	}

	result := dto.TokenResponse{
		Token:       data.Token,
		Type:        "Bearer",
		ExpireHours: h.jwtExpire,
		ExpiredAt:   data.ExpiredAt,
	}

	return response.SendSuccess(c, http.StatusOK, "successfully", result)
}
