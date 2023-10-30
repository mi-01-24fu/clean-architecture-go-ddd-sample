package handler

import (
	"clean-architecture-go-ddd-sample/usecase"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler interface {
	Post() echo.HandlerFunc
}

type userHandler struct {
	signupUsecase usecase.SignupUsecase
}

func NewUserHandler(signupUsecase usecase.SignupUsecase) UserHandler {
	return &userHandler{signupUsecase: signupUsecase}
}

type responseUser struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
}

func (u *userHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {

		var req usecase.SignupRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		signupRequest := usecase.SignupRequest{
			UserName:    req.UserName,
			MailAddress: req.MailAddress,
			Password:    req.Password,
		}

		user, err := u.signupUsecase.Signup(context.Background(), signupRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := usecase.SignupResponse{
			ID:          user.ID,
			UserName:    user.UserName.String(),
			MailAddress: user.MailAddress.String(),
		}

		return c.JSON(http.StatusCreated, res)
	}
}
