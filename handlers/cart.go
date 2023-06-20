package handlers

import (
	"net/http"
	"strconv"
	dtocart "waysbook/dto/cart"
	dtoresult "waysbook/dto/result"
	"waysbook/models"
	"waysbook/repository"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	CartRepository repository.CartRepository
}

func HandlerCart(CartRepository repository.CartRepository) *CartHandler {
	return &CartHandler{CartRepository}
}

func (h *CartHandler) GetAllCart(c echo.Context) error {
	var dataResponse, err = h.CartRepository.GetAllCart()

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
func (h *CartHandler) GetCartById(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var dataResponse, err = h.CartRepository.GetCartById(id)

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

func (h *CartHandler) AddCart(c echo.Context) error {
	request := new(dtocart.AddCart)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	cart := models.Cart{
		UserId: int(userId),
		BookId: request.BookId,
	}

	dataResponse, err := h.CartRepository.AddCart(cart)
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

func (h *CartHandler) DeleteCart(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := int(userLogin.(jwt.MapClaims)["id"].(float64))

	id, _ := strconv.Atoi(c.Param("id"))

	cart, err := h.CartRepository.CartIdForDelete(userId, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	dataResponse, err := h.CartRepository.DeleteCart(cart)

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
