package controllers

import (
	"net/http"
	"taskplanner_golang/helper"
	"taskplanner_golang/models"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password, err := helper.HashPassword(c.FormValue("password"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, status, err := models.Register(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !status {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, status, err := models.Login(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !status {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, res)
}
