package main

import (
	"LearnGoLang2/db"
	"LearnGoLang2/handler"
	"LearnGoLang2/log"
	"LearnGoLang2/repository/repo_impl"
	"LearnGoLang2/router"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func init() {
	os.Setenv("APP_NAME", "")
	log.InitLogger(false)
}

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5433,
		UserName: "postgres",
		Password: "123456",
		DbName:   "LearnGolang2",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	e.Use(middleware.AddTrailingSlash())
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:        e,
		Userhandler: userHandler,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3000"))
}
