package controllers

import (
	"instagram/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	claims := getNilaiToken(c)
	userId := claims.Id
	image, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	caption := c.FormValue("caption")
	createdAt := c.FormValue("createdAt")

	res, err := models.CreatePost(userId, image, caption, createdAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllPost(c echo.Context) error {
	res, err := models.GetAllPost()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetPost(c echo.Context) error {
	id := c.Param("id")
	res, err := models.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func DeletePost(c echo.Context) error {
	id := c.Param("id")
	res, err := models.DeletePost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
