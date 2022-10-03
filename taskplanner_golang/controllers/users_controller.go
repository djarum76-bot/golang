package controllers

import (
	"fmt"
	"net/http"
	"taskplanner_golang/models"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	result, err := models.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetUser(c echo.Context) error {
	claim := getNilaiToken(c)
	id := fmt.Sprintf("%d", claim.Id)
	result, err := models.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
