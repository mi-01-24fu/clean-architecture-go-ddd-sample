package main

import (
	"clean-architecture-go-ddd-sample/crypto"
	"clean-architecture-go-ddd-sample/domain/model"
	"clean-architecture-go-ddd-sample/infrastructure/repository"
	"clean-architecture-go-ddd-sample/interface/handler"
	"clean-architecture-go-ddd-sample/setup"
	"clean-architecture-go-ddd-sample/usecase"
	"github.com/labstack/echo/v4"
)

func main() {

	app := setup.App()
	defer app.DB.Close()

	userFactory := model.NewUserFactory()
	userRepository := repository.NewUserRepository(app.DB, userFactory)
	encrypt := crypto.NewEncrypt()
	signupUsecase := usecase.NewSignupUsecase(userRepository, encrypt)
	userHandler := handler.NewUserHandler(signupUsecase)

	e := echo.New()
	handler.InitRouting(e, userHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
