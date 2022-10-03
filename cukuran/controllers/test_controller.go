package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Pesan":  "Welcome Pantek",
		"Number": 1,
	})
}
