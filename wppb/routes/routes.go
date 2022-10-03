package routes

import (
	"wppb/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Static("/upload", "upload")

	e.POST("/wisata", controllers.AddWisata)
	e.GET("/wisata", controllers.GetAllWisata)

	return e
}
