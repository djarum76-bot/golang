package controllers

import (
	"bloc/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddPost(c echo.Context) error {
	claims := getToken(c)
	user_id := claims.Id
	title := c.FormValue("title")
	body := c.FormValue("body")

	res, err := models.AddPost(user_id, title, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllPost(c echo.Context) error {
	res, err := models.GetAllPost()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetPost(c echo.Context) error {
	id := c.Param("id")
	res, err := models.GetPost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func DeletePost(c echo.Context) error {
	id := c.Param("id")
	res, err := models.DeletePost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdatePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	title := c.FormValue("title")
	body := c.FormValue("body")

	res, err := models.UpdatePost(id, title, body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
