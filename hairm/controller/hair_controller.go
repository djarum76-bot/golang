package controller

import (
	"hairm/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllHair(c echo.Context) error {
	res, err := models.GetAllHair()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetHair(c echo.Context) error {
	id := c.Param("id")

	res, err := models.GetHair(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
