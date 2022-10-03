package controllers

import (
	"crudflutter/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DataBarang(c echo.Context) error {
	result, err := models.DataBarang()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataBarangTertentu(c echo.Context) error {
	id := c.Param("id")
	result, err := models.DataBarangTertentu(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanBarang(c echo.Context) error {
	kodebarang := c.FormValue("kodebarang")
	namabarang := c.FormValue("namabarang")

	result, err := models.SimpanBarang(kodebarang, namabarang)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateBarang(c echo.Context) error {
	id := c.Param("id")

	kodebarang := c.FormValue("kodebarang")
	namabarang := c.FormValue("namabarang")

	result, err := models.UpdateBarang(id, kodebarang, namabarang)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteBarang(c echo.Context) error {
	id := c.Param("id")
	result, err := models.DeleteBarang(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
