package routes

import (
	"waysbook/handlers"
	"waysbook/pkg/middleware"
	"waysbook/pkg/postgresql"
	"waysbook/repository"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Group) {
	bookRepository := repository.RepositoryBook(postgresql.DB)
	h := handlers.HandlerBook(bookRepository)

	e.GET("/books", h.GetAllBook)
	e.GET("/book/:id", h.GetBookById)
	e.POST("/book", middleware.UploadFile(h.AddBook))
	e.DELETE("/book/:id", middleware.Auth(h.DeleteBook))
	e.PATCH("/book/:id", middleware.UpdateBook(h.UpdateBook))
}
