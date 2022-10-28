package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Ryan",
		"email": "ryanLucifeed@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string
		FullName string
	}

	user := User{Email: "ryanLucifeed@gmail.com",
		FullName: "Ryan Locser"}
	return c.JSON(http.StatusOK, user)
}
