package handlers

import (
	"net/http"
	"strconv"
	"time"
	dtobook "waysbook/dto/book"
	dtoresult "waysbook/dto/result"
	"waysbook/models"
	"waysbook/repository"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookRepository repository.BookRepository
}

func HandlerBook(BookRepository repository.BookRepository) *BookHandler {
	return &BookHandler{BookRepository}
}

func (h *BookHandler) GetAllBook(c echo.Context) error {
	var dataResponse, err = h.BookRepository.GetAllBook()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *BookHandler) GetBookById(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	dataResponse, err := h.BookRepository.GetBookById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *BookHandler) AddBook(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	dataImage := c.Get("dataImage").(string)
	publicIdFile := c.Get("filePublicId").(string)
	publicIdImage := c.Get("imagePublicId").(string)

	request := new(dtobook.BookRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	price, _ := strconv.Atoi(c.FormValue("price"))
	discount, _ := strconv.Atoi(c.FormValue("discount"))

	book := models.Book{
		Title:           c.FormValue("title"),
		PublicationDate: c.FormValue("publication_date"),
		Pages:           c.FormValue("pages"),
		ISBN:            c.FormValue("isbn"),
		Author:          c.FormValue("author"),
		Price:           price,
		Discount:        discount,
		DiscountAmount:  price * (100 - discount / 100),
		Description:     c.FormValue("description"),
		BookAttachment:  dataFile,
		Thumbnail:       dataImage,
		PublicIdBook:    publicIdFile,
		PublicIdThumbnail:   publicIdImage,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	dataResponse, err := h.BookRepository.AddBook(book)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	book, err := h.BookRepository.GetBookById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	dataResponse, err := h.BookRepository.DeleteBook(book)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	dataImage := c.Get("dataImage").(string)
	publicIdFile := c.Get("filePublicId").(string)
	publicIdImage := c.Get("imagePublicId").(string)

	request := new(dtobook.BookRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var id, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	book, err := h.BookRepository.GetBookById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	var title = c.FormValue("title")
	if title != "" {
		book.Title = title
	}

	var publicationDate = c.FormValue("publication_date")
	if publicationDate != "" {
		book.PublicationDate = publicationDate
	}

	var pages = c.FormValue("pages")
	if pages != "" {
		book.Pages = pages
	}

	var isbn = c.FormValue("isbn")
	if isbn != "" {
		book.ISBN = isbn
	}

	var author = c.FormValue("author")
	if author != "" {
		book.Author = author
	}

	discount, _ := strconv.Atoi(c.FormValue("discount"))
	if discount != 0 {
		book.Discount = discount
	}

	price, _ := strconv.Atoi(c.FormValue("price"))
	if price != 0 {
		book.Price = price
	}

	// var discountAmount = c.FormValue("discount_amount")
	// if discountAmount != "" {
	// 	book.DiscountAmount = discountAmount
	// }

	var description = c.FormValue("description")
	if description != "" {
		book.Description = description
	}

	if dataFile != "" {
		book.BookAttachment = dataFile
	}

	if dataImage != "" {
		book.Thumbnail = dataImage
	}

	if publicIdFile != "" {
		book.PublicIdBook = publicIdFile
	}

	if publicIdImage != "" {
		book.PublicIdThumbnail = publicIdImage
	}

	book.UpdatedAt = time.Now()

	dataResponse, err := h.BookRepository.UpdateBook(book)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}
