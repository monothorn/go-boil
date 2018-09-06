package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"monothorn/go-boil/middleware"

	cmDelivery "monothorn/go-boil/comments/delivery"
	cmRepository "monothorn/go-boil/comments/repository"
	cmUsecase "monothorn/go-boil/comments/usecase"
	pDelivery "monothorn/go-boil/posts/delivery"
	pRepository "monothorn/go-boil/posts/repository"
	pUseCase "monothorn/go-boil/posts/usecase"

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

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	pR := pRepository.NewPostsRepository("apiv1")
	pUC := pUseCase.NewPostsUseCase(pR, timeoutContext)
	pDelivery.NewPostsHandler(router, pUC)
	cmR := cmRepository.NewCommentsRepository("apiv1")
	cmUC := cmUsecase.NewCommentsUseCase(cmR, timeoutContext)
	cmDelivery.NewCommentsHandler(router, cmUC)
	router.Start(viper.GetString("server.address"))

}
