package controllers

import (
	"net/http"
	"nonolep/models"
	"nonolep/services"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := services.Register(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := services.Login(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetUser(c echo.Context) error {
	claim := getTokenInfo(c)

	res, err := services.GetUser(claim.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateUser(c echo.Context) error {
	claim := getTokenInfo(c)

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	age, err := strconv.Atoi(c.FormValue("age"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	weight, err := strconv.Atoi(c.FormValue("weight"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	height, err := strconv.Atoi(c.FormValue("height"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	user := models.User{
		ID:     claim.ID,
		Gender: c.FormValue("gender"),
		Age:    age,
		Weight: height,
		Height: weight,
		Goals:  form.Value["goals"],
		Level:  c.FormValue("level"),
		Name:   c.FormValue("name"),
		Phone:  c.FormValue("phone"),
	}

	picture, err := c.FormFile("picture")
	if err != nil {
		if err != http.ErrMissingFile {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}
	}

	isEditProfile, err := strconv.ParseBool(c.FormValue("is_edit_profile"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := services.UpdateUser(user, picture, isEditProfile)
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
