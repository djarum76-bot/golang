package controllers

import (
	"net/http"
	"taskplanner_golang/models"

	"github.com/labstack/echo/v4"
)

func AddNote(c echo.Context) error {
	claim := getNilaiToken(c)

	userId := claim.Id
	title := c.FormValue("title")
	content := c.FormValue("content")

	result, err := models.AddNote(userId, title, content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllNote(c echo.Context) error {
	claim := getNilaiToken(c)

	userId := claim.Id

	result, err := models.GetAllNote(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetNote(c echo.Context) error {
	id := c.Param("id")
	result, err := models.GetNote(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteNote(c echo.Context) error {
	claim := getNilaiToken(c)

	id := c.Param("id")
	userId := claim.Id
	result, err := models.DeleteNote(id, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateNote(c echo.Context) error {
	claim := getNilaiToken(c)

	id := c.Param("id")
	userId := claim.Id
	title := c.FormValue("title")
	content := c.FormValue("content")

	result, err := models.UpdateNote(id, userId, title, content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
