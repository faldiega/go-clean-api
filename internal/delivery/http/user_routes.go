package http

import "github.com/labstack/echo/v4"

func RegisterUserRoutes(v1Group *echo.Group, handler *UserHandler) {
	v1Group.GET("/users", handler.GetUserList)
	v1Group.GET("/users/:id", handler.GetUser)
	v1Group.POST("/users", handler.CreateUser)
	v1Group.PUT("/users/:id", handler.UpdateUser)
	v1Group.DELETE("/users/:id", handler.DeleteUser)
}
