package controllers

import (
	"instagram/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateComment(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	claims := getNilaiToken(c)
	userId := claims.Id
	comment := c.FormValue("comment")
	createdAt := c.FormValue("createdAt")

	res, err := models.CreateComment(postId, userId, comment, createdAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllComment(c echo.Context) error {
	postId := c.Param("post_id")
	res, err := models.GetAllComment(postId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
