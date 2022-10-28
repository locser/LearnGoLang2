package main

import (
	"LearnGoLang2/db"
	"LearnGoLang2/handler"
	"LearnGoLang2/log"
	"context"
	"github.com/labstack/echo/v4"
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

	log.Error("Co loi xay ra")

	sql.Connect()
	defer sql.Close()

	var email string
	err := sql.Db.GetContext(context.Background(), &email, "select email from users where  email= $1", "abc@gmail.com")
	if err != nil {
		log.Error(err.Error())
	}
	e := echo.New()
	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)
	e.Logger.Fatal(e.Start(":3000"))
}
