package routes

import (
	"bloc_socket/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", controllers.GetAllUser)
	e.POST("/user", controllers.AddUser)

	return e
}
