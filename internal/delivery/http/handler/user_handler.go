package handler

import (
	"go-simple-api/internal/config"
	"go-simple-api/internal/delivery/http/dto"
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/usecase"
	"go-simple-api/internal/initialize"
	"go-simple-api/internal/utils"
	pkgLogger "go-simple-api/pkg/common/logger"
	"go-simple-api/pkg/common/pagination"
	"go-simple-api/pkg/common/response"
	customValidator "go-simple-api/pkg/common/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserHandler struct {
	usecase usecase.UserUsecase
	config  *config.Config
	logger  *zap.Logger
}

func NewUserHandler(uc usecase.UserUsecase, container *initialize.Container) *UserHandler {
	return &UserHandler{
		usecase: uc,
		config:  container.Config,
		logger:  container.ZapLogger,
	}
}

func (h *UserHandler) GetUserList(c echo.Context) error {

	defaultPage := h.config.Pagination.DefaultPage
	defaultLimit := h.config.Pagination.DefaultLimit

	// ambil logger dengan trace ID dari context
	logger := pkgLogger.FromContext(c.Request().Context(), h.logger)

	// 1. ambil pagination dari query param
	p := pagination.GetPaginationFromCtx(c, defaultPage, defaultLimit)

	// 2. ambil data dari usecase
	users, totalItems, err := h.usecase.GetUsers(c.Request().Context(), p)
	if err != nil {
		logger.Error("failed to fetch users", zap.Error(err))
		return response.SendError(c, http.StatusInternalServerError, "failed to fetch users", err.Error())
	}

	var data []dto.UserResponse

	for _, u := range users {
		data = append(data, dto.UserResponse{
			ID:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			IsActive:    u.IsActive,
			CreatedDate: *utils.ToDatetime(u.CreatedDate),
			UpdatedDate: utils.ToDatetime(u.UpdatedDate),
		})
	}

	logger.Info("users fetched successfully", zap.Int("total", totalItems))

	// 3. build hasil paging
	result := pagination.BuildResult(data, p, totalItems)

	return response.SendSuccess(c, http.StatusOK, "successfully", result)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SendError(c, http.StatusBadRequest, "bad request param ID", err.Error())
	}

	result, err := h.usecase.GetUser(c.Request().Context(), int(id))
	if err != nil {
		return response.SendError(c, http.StatusNotFound, "user not found", err.Error())
	}

	data := dto.UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		IsActive:    result.IsActive,
		CreatedDate: *utils.ToDatetime(result.CreatedDate),
		UpdatedDate: utils.ToDatetime(result.UpdatedDate),
	}

	return response.SendSuccess(c, http.StatusOK, "successfully", data)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	req := new(dto.CreateUserRequest)
	if err := c.Bind(&req); err != nil {
		return response.SendError(c, http.StatusBadRequest, "invalid request", err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.SendError(c, http.StatusBadRequest, "validation failed", customValidator.FormatValidationError(err))
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		IsActive: true,
	}

	result, err := h.usecase.CreateUser(c.Request().Context(), user)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "create user failed", err.Error())

	}

	data := dto.UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		IsActive:    result.IsActive,
		CreatedDate: *utils.ToDatetime(result.CreatedDate),
	}

	return response.SendSuccess(c, http.StatusOK, "successfully", data)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SendError(c, http.StatusBadRequest, "invalid ID", err.Error())

	}

	req := new(dto.UpdateUserRequest)
	if err := c.Bind(&req); err != nil {
		return response.SendError(c, http.StatusBadRequest, "invalid request", err.Error())
	}

	if err := c.Validate(req); err != nil {
		return response.SendError(c, http.StatusBadRequest, "validation failed", customValidator.FormatValidationError(err))
	}

	user := entity.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		IsActive: req.IsActive,
	}

	result, err := h.usecase.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "update user failed", err.Error())
	}

	data := dto.UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		IsActive:    result.IsActive,
		CreatedDate: *utils.ToDatetime(result.CreatedDate),
		UpdatedDate: utils.ToDatetime(result.UpdatedDate),
	}

	return response.SendSuccess(c, http.StatusOK, "successfully", data)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SendError(c, http.StatusBadRequest, "invalid ID", err.Error())
	}

	user, err := h.usecase.GetUser(c.Request().Context(), int(id))
	if err != nil {
		return response.SendError(c, http.StatusNotFound, "user not found", err.Error())
	}

	err = h.usecase.DeleteUser(c.Request().Context(), int(id))

	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "delete user failed", err.Error())
	}

	return response.SendSuccess(c, http.StatusOK, "successfully", "user ["+user.Name+"] has been deleted.")
}
