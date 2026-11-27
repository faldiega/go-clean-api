package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"go-clean-api/internal/delivery/http/dto"
	"go-clean-api/internal/domain/entity"
	"go-clean-api/internal/domain/usecase"
)

type HealthHandler struct {
	usecase usecase.HealthUsecase
}

func NewHealthHandler(uc usecase.HealthUsecase) *HealthHandler {
	return &HealthHandler{usecase: uc}
}

// GET /health/live
func (h *HealthHandler) Liveness(c echo.Context) error {
	result := h.usecase.Liveness(c.Request().Context())
	return c.JSON(http.StatusOK, dto.ToHealthResponse(result))
}

// GET /health/ready
func (h *HealthHandler) Readiness(c echo.Context) error {
	result := h.usecase.Readiness(c.Request().Context())

	code := http.StatusOK
	if result.Status != entity.HealthUp {
		code = http.StatusServiceUnavailable
	}
	return c.JSON(code, dto.ToHealthResponse(result))
}
