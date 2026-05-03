package handler

import (
	"go-simple-api/internal/delivery/http/dto"
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/usecase"
	"go-simple-api/pkg/common/response"
	customValidator "go-simple-api/pkg/common/validator"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc}
}

func (h *UserHandler) GetUserList(c echo.Context) error {
	users, err := h.usecase.GetUsers()
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "failed to fetch users", err.Error())
	}

	var data []dto.UserResponse

	for _, u := range users {
		data = append(data, dto.UserResponse{
			ID:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			IsActive:    u.IsActive,
			CreatedDate: u.CreatedDate.Format(time.DateTime),
			UpdatedDate: func() *string {
				if u.UpdatedDate != nil {
					date := u.UpdatedDate.Format(time.DateTime)
					return &date
				}

				return nil
			}(),
		})
	}

	return response.SendSuccess(c, http.StatusOK, "successfully", data)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SendError(c, http.StatusBadRequest, "bad request param ID", err.Error())
	}

	result, err := h.usecase.GetUser(int(id))
	if err != nil {
		return response.SendError(c, http.StatusNotFound, "user not found", err.Error())
	}

	data := dto.UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		IsActive:    result.IsActive,
		CreatedDate: result.CreatedDate.Format(time.DateTime),
		UpdatedDate: func() *string {
			date := result.UpdatedDate.Format(time.DateTime)
			return &date
		}(),
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

	result, err := h.usecase.CreateUser(user)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "create user failed", err.Error())

	}

	data := dto.UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		IsActive:    result.IsActive,
		CreatedDate: result.CreatedDate.Format(time.DateTime),
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

	result, err := h.usecase.UpdateUser(user)
	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "update user failed", err.Error())

	}

	data := dto.UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		IsActive:    result.IsActive,
		CreatedDate: result.CreatedDate.Format(time.DateTime),
		UpdatedDate: func() *string {
			date := result.UpdatedDate.Format(time.DateTime)
			return &date
		}()}

	return response.SendSuccess(c, http.StatusOK, "successfully", data)

}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SendError(c, http.StatusBadRequest, "invalid ID", err.Error())
	}

	user, err := h.usecase.GetUser(int(id))
	if err != nil {
		return response.SendError(c, http.StatusNotFound, "user not found", err.Error())
	}

	err = h.usecase.DeleteUser(int(id))

	if err != nil {
		return response.SendError(c, http.StatusInternalServerError, "delete user failed", err.Error())
	}

	return response.SendSuccess(c, http.StatusOK, "successfully", "user ["+user.Name+"] has been deleted.")
}
