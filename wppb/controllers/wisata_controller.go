package controllers

import (
	"net/http"
	"wppb/models"

	"github.com/labstack/echo/v4"
)

func AddWisata(c echo.Context) error {
	name := c.FormValue("name")
	location := c.FormValue("location")
	description := c.FormValue("description")
	openDay := c.FormValue("openDay")
	openTime := c.FormValue("openTime")
	ticketPrice := c.FormValue("ticketPrice")
	image, err := c.FormFile("image")
	if err != nil {
		return err
	}
	result, err := models.AddWisata(name, location, description, openDay, openTime, ticketPrice, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllWisata(c echo.Context) error {
	result, err := models.GetAllWisata()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
