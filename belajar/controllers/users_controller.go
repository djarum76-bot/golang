package controllers

import (
	"belajar/helper"
	"belajar/models"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
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

	res, err := models.Register(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
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

func GetAllUser(c echo.Context) error {
	result, err := models.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CetakToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"username": claims.Username,
		"id":       claims.Id,
	})
}

func DeleteUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)
	id := fmt.Sprintf("%d", claims.Id)
	result, err := models.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func Upload(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("upload/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "berhasil")
}
