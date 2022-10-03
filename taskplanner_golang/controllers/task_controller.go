package controllers

import (
	"net/http"
	"taskplanner_golang/models"

	"github.com/labstack/echo/v4"
)

func AddTask(c echo.Context) error {
	claim := getNilaiToken(c)

	userId := claim.Id
	title := c.FormValue("title")
	tanggal := c.FormValue("tanggal")
	waktu := c.FormValue("waktu")
	date := c.FormValue("date")
	deskripsi := c.FormValue("deskripsi")

	result, err := models.AddTask(userId, title, tanggal, waktu, date, deskripsi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTask(c echo.Context) error {
	claim := getNilaiToken(c)

	userId := claim.Id

	result, err := models.GetAllTask(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTaskDate(c echo.Context) error {
	claim := getNilaiToken(c)

	userId := claim.Id

	result, err := models.GetAllTaskDate(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTask(c echo.Context) error {
	id := c.Param("id")
	result, err := models.GetTask(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteTask(c echo.Context) error {
	claim := getNilaiToken(c)

	userId := claim.Id
	id := c.Param("id")

	result, err := models.DeleteTask(id, userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateTask(c echo.Context) error {
	claim := getNilaiToken(c)

	id := c.Param("id")
	userId := claim.Id
	title := c.FormValue("title")
	tanggal := c.FormValue("tanggal")
	waktu := c.FormValue("waktu")
	date := c.FormValue("date")
	deskripsi := c.FormValue("deskripsi")

	result, err := models.UpdateTask(id, userId, title, tanggal, waktu, date, deskripsi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
