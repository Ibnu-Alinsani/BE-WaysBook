package handlers

import (
	"log"
	"net/http"
	"time"
	authdto "waysbook/dto/auth"
	dtoresult "waysbook/dto/result"
	userdto "waysbook/dto/user"
	"waysbook/models"
	"waysbook/pkg/bcrypt"
	jwtToken "waysbook/pkg/jwt"
	"waysbook/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthRepository repository.AuthRepository
}

func HandlerAuth(AuthRepository repository.AuthRepository) *AuthHandler {
	return &AuthHandler{AuthRepository}
}

func (h *AuthHandler) Register(c echo.Context) error {
	request := new(authdto.RegisterRequest)
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
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	password, err := bcrypt.HashingPassword(request.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	user := models.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  password,
		Gender:    request.Gender,
		Phone:     request.Phone,
		Address:   request.Address,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dataResponse, err := h.AuthRepository.Register(user)

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

func (h *AuthHandler) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	checkUser, err := h.AuthRepository.Login(user.Email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	isValid := bcrypt.CheckPasswordHash(checkUser.Password, user.Password)

	if !isValid {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "Wrong Password",
		})
	}

	claims := jwt.MapClaims{}
	claims["id"] = checkUser.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := jwtToken.GenerateToken(&claims)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse{
		Name:  checkUser.Name,
		Email: checkUser.Email,
		Gender: checkUser.Gender,
		Phone: checkUser.Phone,
		Address: checkUser.Address,
		Avatar: checkUser.Avatar,
		Role:  checkUser.Role,
		Token: token,
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: loginResponse,
	})
}

func (h *AuthHandler) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: user,
	})
}

func ConvertResponseUser(user models.User) userdto.UserResponse {
	return userdto.UserResponse{
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Gender:  user.Gender,
		Avatar:  user.Avatar,
		Address: user.Address,
		Role:    user.Role,
	}
}
