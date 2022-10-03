package routes

import (
	"crud_stream/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/user", controllers.AddUser)
	e.GET("/stream", controllers.RealTime)

	return e
}
