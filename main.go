package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"monothorn/go-boil/middleware"
	posts "monothorn/go-boil/posts/delivery"
	_postsRepo "monothorn/go-boil/posts/repository"
	postsUseCase "monothorn/go-boil/posts/usecase"
	"monothorn/go-boil/utilities"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}
func main() {
	router := echo.New()
	middL := middleware.InitMiddleware()
	router.Use(middL.CORS)
	cmR := _postsRepo.NewPostsRepository("apiv1")
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	cmUC := postsUseCase.NewPostsUseCase(cmR, timeoutContext)
	posts.NewPostsHandler(router, cmUC)
	router.GET("/test", func(c echo.Context) error {
		db, _ := utilities.GetInstance(viper.GetString("sql.dsn"))
		rows, err := db.Query("SELECT id FROM test_db.users")
		for rows.Next() {
			var uid int
			err = rows.Scan(&uid)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(uid)
		}
		return nil

	})
	router.Start(viper.GetString("server.address"))

}
