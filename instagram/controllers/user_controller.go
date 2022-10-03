package controllers

import (
	"instagram/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CompleteProfile(c echo.Context) error {
	claims := getNilaiToken(c)
	id := claims.Id
	phone := c.FormValue("phone")
	picture, err := c.FormFile("picture")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	res, err := models.CompleteProfile(id, phone, picture)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetUser(c echo.Context) error {
	claims := getNilaiToken(c)
	id := claims.Id

	res, err := models.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
