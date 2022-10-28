package main

import (
	"LearnGoLang2/db"
	"LearnGoLang2/handler"
	"github.com/labstack/echo/v4"
	_ "github.com/rubenv/sql-migrate"
)

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
	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)
	e.Logger.Fatal(e.Start(":3000"))
}
