package routes

import (
	"waysbook/handlers"
	"waysbook/pkg/middleware"
	"waysbook/pkg/postgresql"
	"waysbook/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repository.RepositoryUser(postgresql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.GetAllUser)
	e.GET("/user/:id", h.GetUserById)
	e.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
	e.PATCH("/user/:id", middleware.UpdateUser(h.UpdateUser))
}