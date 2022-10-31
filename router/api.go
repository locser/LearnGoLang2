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
	// user
	api.Echo.POST("/user/sign-in", api.Userhandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)

	// profile
	user := api.Echo.Group("/user", middleware2.JWTMiddleware())
	user.GET("/profile", api.Userhandler.HandleProfile)
	user.PUT("/profile/update", api.Userhandler.UpdateProfile)
}
