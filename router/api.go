package router

import (
	"LearnGoLang2/handler"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	Userhandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.GET("/", handler.Welcome)
	api.Echo.POST("/user/sign-in", api.Userhandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)
}
