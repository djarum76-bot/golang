package controllers

import (
	"net/http"
	"os"
	"yusril/helper"
	"yusril/models"
	"yusril/services"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	email, err := helper.Encrypt(c.FormValue("email"), os.Getenv("SECRET_STRING"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	password, err := helper.HashPassword(c.FormValue("password"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := services.Register(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	email, err := helper.Encrypt(c.FormValue("email"), os.Getenv("SECRET_STRING"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	password := c.FormValue("password")

	res, err := services.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func getTokenInfo(c echo.Context) *models.JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	return claims
}
