package controllers

import (
	"instagram/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LikePost(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	claims := getNilaiToken(c)
	userId := claims.Id

	res, err := models.LikePost(postId, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func UnlikePost(c echo.Context) error {
	id := c.Param("id")

	res, err := models.UnlikePost(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
