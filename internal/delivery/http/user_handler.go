package http

import (
	"go-simple-api/internal/domain/entity"
	"go-simple-api/internal/domain/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(e *echo.Echo, uc usecase.UserUsecase) {

	handler := &UserHandler{uc}

	e.GET("/users", handler.GetUserList)
	e.GET("/users/:id", handler.GetUser)
	e.POST("/users", handler.CreateUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
}

func (h *UserHandler) GetUserList(c echo.Context) error {

	users, err := h.usecase.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.usecase.GetUser(uint(id))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {

	var user entity.User

	c.Bind(&user)

	result, err := h.usecase.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	var user entity.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request body")
	}

	user.ID = uint(id)

	result, err := h.usecase.UpdateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	user, err := h.usecase.GetUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	err = h.usecase.DeleteUser(uint(id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user deleted. [" + user.Name + "]",
	})
}
