package routes

import (
	"hairm/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/hairs", controller.GetAllHair)
	e.GET("/hair/:id", controller.GetHair)

	return e
}
