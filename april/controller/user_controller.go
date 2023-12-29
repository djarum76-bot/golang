package controller

import (
	"april/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, message, err := service.Register(name, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": message,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, message, err := service.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": message,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetUser(c echo.Context) error {
	user := getTokenInfo(c)
	ID := user.ID

	res, err := service.GetUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
