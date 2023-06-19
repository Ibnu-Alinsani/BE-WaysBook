package routes

import (
	"waysbook/handlers"
	"waysbook/pkg/middleware"
	"waysbook/pkg/postgresql"
	"waysbook/repository"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transaction := repository.RepositoryTransaction(postgresql.DB)
	h := handlers.HandlerTransaction(transaction)

	e.GET("/transactions", middleware.Auth(h.GetAllTransaction))
	e.GET("/transaction/:id", middleware.Auth(h.GetTransactionById))
	e.POST("/transaction", middleware.Auth(h.AddTransaction))
	e.POST("/notification", h.Notification)
}
