package routes

import (
	"jwt/controllers"
	"jwt/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]string

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, M{"Ikan": "Paus"})
	})

	e.GET("/pegawai", controllers.GetAllPegawai, middleware.IsAuthenticated)
	e.GET("/pegawai/:id", controllers.GetPegawai)
	e.POST("/pegawai", controllers.AddPegawai)
	e.PUT("/pegawai/:id", controllers.UpdatePegawai)
	e.DELETE("/pegawai/:id", controllers.DeletePegawai)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	return e
}
