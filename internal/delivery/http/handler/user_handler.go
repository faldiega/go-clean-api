package handler

import (
	"go-simple-api/internal/delivery/http/dto"
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/usecase"
	"net/http"
	"strconv"

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
		return c.JSON(http.StatusInternalServerError, err)
	}

	var response []dto.UserResponse

	for _, u := range users {
		response = append(response, dto.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	user, err := h.usecase.GetUser(int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	response := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var req dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		IsActive: true,
	}

	result, err := h.usecase.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := dto.UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		IsActive:    result.IsActive,
		CreatedDate: result.CreatedDate.String(),
		UpdatedDate: result.UpdatedDate.String(),
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	var req dto.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	user := entity.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		IsActive: req.IsActive,
	}

	result, err := h.usecase.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	response := dto.UserResponse{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		IsActive:    result.IsActive,
		CreatedDate: result.CreatedDate.String(),
		UpdatedDate: result.UpdatedDate.String(),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	user, err := h.usecase.GetUser(int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	err = h.usecase.DeleteUser(int(id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user deleted. [" + user.Name + "]",
	})
}
