package routes

import (
	"waysbook/handlers"
	"waysbook/pkg/middleware"
	"waysbook/pkg/postgresql"
	"waysbook/repository"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	cart := repository.RepositoryCart(postgresql.DB)
	h := handlers.HandlerCart(cart)

	e.GET("/carts", h.GetAllCart)
	e.GET("/cart/:id", h.GetCartById)
	e.DELETE("/cart/:id", middleware.Auth(h.DeleteCart))
	e.POST("/cart", middleware.Auth(h.AddCart))
}