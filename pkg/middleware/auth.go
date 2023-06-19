package middleware

import (
	"net/http"
	"strings"
	dtoresult "waysbook/dto/result"
	jwtToken "waysbook/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dtoresult.ErrorResult{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized, login dulu mang",
			})
		}

		
		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, dtoresult.ErrorResult{
				Code:    http.StatusUnauthorized,
				Message: token,
			})
		}

		c.Set("userLogin", claims)
		return next(c)
	}

}
