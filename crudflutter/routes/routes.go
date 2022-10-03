package routes

import (
	"crudflutter/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hi")
	})

	e.GET("/barang", controllers.DataBarang)
	e.GET("/barang/:id", controllers.DataBarangTertentu)
	e.POST("/barang", controllers.SimpanBarang)
	e.PUT("/barang/:id", controllers.UpdateBarang)
	e.DELETE("/barang/:id", controllers.DeleteBarang)

	return e
}
