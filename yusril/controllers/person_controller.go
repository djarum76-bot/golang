package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"yusril/helper"
	"yusril/models"
	"yusril/services"

	"github.com/labstack/echo/v4"
)

func AddPersonWithImage(c echo.Context) error {
	claim := getTokenInfo(c)

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	email, err := helper.EncryptList(form.Value["email"])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	person := models.Person{
		ID:               0,
		UserID:           claim.ID,
		Nickname:         c.FormValue("nickname"),
		FirstName:        c.FormValue("first_name"),
		LastName:         c.FormValue("last_name"),
		Email:            email,
		Job:              c.FormValue("job"),
		Phone:            c.FormValue("phone"),
		DateOfBirth:      c.FormValue("date_of_birth"),
		PlaceOfBirth:     c.FormValue("place_of_birth"),
		PlaceOfResidence: form.Value["place_of_residence"],
		Instagram:        c.FormValue("instagram"),
		Twitter:          c.FormValue("twitter"),
		Facebook:         c.FormValue("facebook"),
		LinkedIn:         c.FormValue("linkedin"),
		Major:            form.Value["major"],
		College:          form.Value["college"],
		Organization:     form.Value["organization"],
		Party:            form.Value["party"],
		LastActivity:     form.Value["last_activity"],
		FirstMeet:        c.FormValue("first_meet"),
		Category:         c.FormValue("category"),
		Tags:             form.Value["tags"],
		CreatedAt:        c.FormValue("created_at"),
	}

	image, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	images := form.File["images"]

	res, err := services.AddPersonWithImage(person, image, images)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func AddPersonWithoutImage(c echo.Context) error {
	claim := getTokenInfo(c)

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	email, err := helper.EncryptList(form.Value["email"])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	person := models.Person{
		ID:               0,
		UserID:           claim.ID,
		Nickname:         c.FormValue("nickname"),
		FirstName:        c.FormValue("first_name"),
		LastName:         c.FormValue("last_name"),
		Email:            email,
		Job:              c.FormValue("job"),
		Phone:            c.FormValue("phone"),
		DateOfBirth:      c.FormValue("date_of_birth"),
		PlaceOfBirth:     c.FormValue("place_of_birth"),
		PlaceOfResidence: form.Value["place_of_residence"],
		Instagram:        c.FormValue("instagram"),
		Twitter:          c.FormValue("twitter"),
		Facebook:         c.FormValue("facebook"),
		LinkedIn:         c.FormValue("linkedin"),
		Major:            form.Value["major"],
		College:          form.Value["college"],
		Organization:     form.Value["organization"],
		Party:            form.Value["party"],
		LastActivity:     form.Value["last_activity"],
		FirstMeet:        c.FormValue("first_meet"),
		Category:         c.FormValue("category"),
		Tags:             form.Value["tags"],
		CreatedAt:        c.FormValue("created_at"),
	}

	images := form.File["images"]

	res, err := services.AddPersonWithoutImage(person, images)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllPerson(c echo.Context) error {
	claim := getTokenInfo(c)

	res, err := services.GetAllPerson(claim.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetPerson(c echo.Context) error {
	claim := getTokenInfo(c)
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := services.GetPerson(claim.ID, ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdatePerson(c echo.Context) error {
	claim := getTokenInfo(c)

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	email, err := helper.EncryptList(form.Value["email"])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	person := models.Person{
		ID:               id,
		UserID:           claim.ID,
		Nickname:         c.FormValue("nickname"),
		FirstName:        c.FormValue("first_name"),
		LastName:         c.FormValue("last_name"),
		Email:            email,
		Job:              c.FormValue("job"),
		Phone:            c.FormValue("phone"),
		DateOfBirth:      c.FormValue("date_of_birth"),
		PlaceOfBirth:     c.FormValue("place_of_birth"),
		PlaceOfResidence: form.Value["place_of_residence"],
		Instagram:        c.FormValue("instagram"),
		Twitter:          c.FormValue("twitter"),
		Facebook:         c.FormValue("facebook"),
		LinkedIn:         c.FormValue("linkedin"),
		Major:            form.Value["major"],
		College:          form.Value["college"],
		Organization:     form.Value["organization"],
		Party:            form.Value["party"],
		LastActivity:     form.Value["last_activity"],
		FirstMeet:        c.FormValue("first_meet"),
		Category:         c.FormValue("category"),
		Tags:             form.Value["tags"],
	}

	res, err := services.UpdatePerson(person)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateImage(c echo.Context) error {
	image, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := services.UpdateImage(image, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateImages(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	images := form.File["images"]

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := services.UpdateImages(images, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func DeletePerson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := services.DeletePerson(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func SearchPerson(c echo.Context) error {
	claim := getTokenInfo(c)
	query := strings.ToLower(c.QueryParam("query"))
	category := strings.ToLower(c.QueryParam("category"))
	tag := strings.ToLower(c.QueryParam("tag"))
	res, err := services.SearchPerson(query, category, tag, claim.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllCategory(c echo.Context) error {
	claim := getTokenInfo(c)
	res, err := services.GetAllCategory(claim.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllTags(c echo.Context) error {
	claim := getTokenInfo(c)
	res, err := services.GetAllTags(claim.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
