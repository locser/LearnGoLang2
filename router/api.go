package router

import (
	"LearnGoLang2/handler"
	middleware2 "LearnGoLang2/middleware"
	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	Userhandler handler.UserHandler
	RepoHandler handler.RepoHandler
}

func (api *API) SetupRouter() {
	// user
	api.Echo.POST("/user/sign-in", api.Userhandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.Userhandler.HandleSignUp)

	// profile
	user := api.Echo.Group("/user", middleware2.JWTMiddleware())
	user.GET("/profile", api.Userhandler.HandleProfile)
	user.PUT("/profile/update", api.Userhandler.UpdateProfile)

	// github repo
	github := api.Echo.Group("/github", middleware2.JWTMiddleware())
	github.GET("/trending", api.RepoHandler.RepoTrending)

	// bookmark
	bookmark := api.Echo.Group("/bookmark", middleware2.JWTMiddleware())
	bookmark.GET("/list", api.RepoHandler.SelectBookmarks)
	bookmark.POST("/add", api.RepoHandler.Bookmark)
	bookmark.DELETE("/delete", api.RepoHandler.DelBookmark)

}
