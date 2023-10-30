package handler

import (
	"clean-architecture-go-ddd-sample/usecase"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler interface {
	Signup() echo.HandlerFunc
}

type userHandler struct {
	signupUsecase usecase.SignupUsecase
}

func NewUserHandler(signupUsecase usecase.SignupUsecase) UserHandler {
	return &userHandler{signupUsecase: signupUsecase}
}

func (u *userHandler) Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		var signupRequest usecase.SignupRequest
		// bind(変換・代入)
		// バインド処理はSignupRequest構造体のタグ部分（json:"user_name"など）に
		//基づいて行われる
		//このタグは、JSONのキー名とGoのフィールド名とのマッピングを定義している
		if err := c.Bind(&signupRequest); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
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
