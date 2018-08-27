package main

import (
	"fmt"
	"time"

	"monothorn/go-boil/middleware"
	posts "monothorn/go-boil/posts/delivery"
	_postsRepo "monothorn/go-boil/posts/repository"
	postsUseCase "monothorn/go-boil/posts/usecase"

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
	router.Start(viper.GetString("server.address"))

}
