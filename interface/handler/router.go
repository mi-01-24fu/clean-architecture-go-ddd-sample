package handler

import "github.com/labstack/echo/v4"

func InitRouting(e *echo.Echo, userHandler UserHandler) {
	e.POST("/signup", userHandler.Signup())
}
