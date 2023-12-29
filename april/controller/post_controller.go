package controller

import (
	"april/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	user := getTokenInfo(c)
	userID := user.ID
	post := c.FormValue("post")
	createdAt := c.FormValue("created_at")

	err := service.CreatePost(userID, post, createdAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Create Post Successfull",
	})
}

func GetAllPost(c echo.Context) error {
	user := getTokenInfo(c)
	userID := user.ID

	res, err := service.GetAllPost(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetPost(c echo.Context) error {
	user := getTokenInfo(c)
	userID := user.ID
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := service.GetPost(userID, ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdatePost(c echo.Context) error {
	user := getTokenInfo(c)
	userID := user.ID
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	post := c.FormValue("post")

	err = service.UpdatePost(userID, ID, post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Update Post Successfull",
	})
}

func DeletePost(c echo.Context) error {
	user := getTokenInfo(c)
	userID := user.ID
	ID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	err = service.DeletePost(userID, ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Delete Post Successfull",
	})
}
