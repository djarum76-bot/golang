package route

import (
	"mei/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)

	return e
}
