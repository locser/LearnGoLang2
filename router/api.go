package router

import (
	"LearnGoLang2/handler"
	middleware2 "LearnGoLang2/middleware"
	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	Userhandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.Userhandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)
	api.Echo.POST("/user/profile", api.Userhandler.HandleProfile, middleware2.JWTMiddleware())

}
