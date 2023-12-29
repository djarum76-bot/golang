package controller

import (
	"mei/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := service.Register(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(res.Status, res)
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := service.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(res.Status, res)
}
