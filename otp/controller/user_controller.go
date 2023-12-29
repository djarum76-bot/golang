package controller

import (
	"net/http"
	"otp/model"
	"otp/service"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	phone := c.FormValue("phone")
	password := c.FormValue("password")

	res, err := service.Register(phone, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	phone := c.FormValue("phone")
	password := c.FormValue("password")

	res, err := service.Login(phone, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func SendOTP(c echo.Context) error {
	user := getTokenInfo(c)
	ID := user.ID
	phone := user.Phone

	res, err := service.SendOTP(ID, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func CheckOTP(c echo.Context) error {
	otp, err := strconv.Atoi(c.FormValue("otp"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	user := getTokenInfo(c)
	ID := user.ID

	res, err := service.CheckOTP(otp, ID)
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
