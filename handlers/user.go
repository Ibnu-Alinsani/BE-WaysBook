package handlers

import (
	"net/http"
	"strconv"
	"time"
	dto "waysbook/dto/result"
	"waysbook/repository"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{
	UserRepository repository.UserRepository
}

func HandlerUser(UserRepository repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository}
}

func (h *UserHandler) GetAllUser(c echo.Context) error {
	dataResponse, err := h.UserRepository.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *UserHandler) GetUserById(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	dataResponse, err := h.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	dataResponse, err := h.UserRepository.DeleteUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	dataFile := c.Get("imageUser").(string)
	publicId := c.Get("imageUserPublicId").(string)

	var id, _ = strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var name = c.FormValue("name")
	if name != "" {
		user.Name = name
	}

	var email = c.FormValue("email")
	if email != "" {
		user.Email = email
	}

	var phone = c.FormValue("phone")
	if phone != "" {
		user.Phone = phone
	}

	var address = c.FormValue("address")
	if address != "" {
		user.Address = address
	}

	if dataFile != "" {
		user.Avatar = dataFile
	}

	if publicId != "" {
		user.PublicIDAvatar = publicId
	}

	user.UpdatedAt = time.Now()

	dataResponse, err := h.UserRepository.UpdateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}