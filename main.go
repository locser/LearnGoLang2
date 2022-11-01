package main

import (
	"LearnGoLang2/db"
	"LearnGoLang2/handler"
	"LearnGoLang2/helper"
	"LearnGoLang2/log"
	"LearnGoLang2/repository/repo_impl"
	"LearnGoLang2/router"
	"fmt"
	"github.com/labstack/echo"
	"os"
	"time"
)

func init() {
	fmt.Println(">>>>", os.Getenv("APP_NAME"))
	//os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {

	sql := &db.Sql{
		Host:     "host.docker.internal", //,localhost, 192.168.1.163, host.docker.internal
		Port:     5433,
		UserName: "postgres",
		Password: "123456",
		DbName:   "LearnGolang2",
	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}

	api := router.API{
		Echo:        e,
		Userhandler: userHandler,
		RepoHandler: repoHandler,
	}

	api.SetupRouter()

	go scheduleUpdateTrending(360*time.Second, repoHandler)

	e.Logger.Fatal(e.Start(":3000"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}
