package controller

import (
	"latihan/helper"
	"latihan/model"
	"latihan/service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	email := c.FormValue("email")
	password, err := helper.HashPassword(c.FormValue("password"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	res, err := service.Register(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := service.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllUser(c echo.Context) error {
	user := getTokenInfo(c)
	ID := user.ID

	res, err := service.GetAllUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetUserProfile(c echo.Context) error {
	user := getTokenInfo(c)
	ID := user.ID

	res, err := service.GetUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetUser(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	res, err := service.GetUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func DeleteUser(c echo.Context) error {
	user := getTokenInfo(c)
	ID := user.ID

	res, err := service.DeleteUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateUser(c echo.Context) error {
	user := getTokenInfo(c)
	ID := user.ID

	email := c.FormValue("email")

	res, err := service.UpdateUser(ID, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func getTokenInfo(c echo.Context) *model.JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)

	return claims
}
