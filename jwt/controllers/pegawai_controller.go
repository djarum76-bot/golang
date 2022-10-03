package controllers

import (
	"jwt/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]string

func GetAllPegawai(c echo.Context) error {
	result, err := models.GetAllPegawai()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetPegawai(c echo.Context) error {
	id := c.Param("id")

	result, err := models.GetPegawai(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddPegawai(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.AddPegawai(nama, alamat, telepon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePegawai(c echo.Context) error {
	id := c.Param("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.UpdatePegawai(id, nama, alamat, telepon)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeletePegawai(c echo.Context) error {
	id := c.Param("id")

	result, err := models.DeletePegawai(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, M{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
