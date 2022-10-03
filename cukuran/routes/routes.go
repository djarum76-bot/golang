package routes

import (
	"github.com/djarum76-bot/cukuran/controllers"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/test", controllers.Test)

	return e
}
