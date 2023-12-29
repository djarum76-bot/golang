package controllers

import (
	"net/http"
	"nonolep/helpers"
	"nonolep/models"
	"nonolep/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddWorkout(c echo.Context) error {
	duration, err := strconv.Atoi(c.FormValue("duration"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	workout := models.Workout{
		Name:     c.FormValue("name"),
		Duration: duration,
	}

	picture, err := c.FormFile("picture")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := services.AddWorkout(workout, picture)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func AddWorkoutDetail(c echo.Context) error {
	var workoutDetail models.WorkoutDetail
	arrWorkoutDetail := []models.WorkoutDetail{}

	workouts, err := services.GetAllWorkout()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	for _, workout := range workouts.Data.([]models.Workout) {
		for i := 0; i < 16; i++ {
			name, duration, picture := helpers.GetDummyData(i)

			workoutDetail = models.WorkoutDetail{
				WorkoutID: workout.ID,
				Name:      name,
				Duration:  duration,
				Picture:   picture,
			}

			arrWorkoutDetail = append(arrWorkoutDetail, workoutDetail)
		}
	}

	res, err := services.AddWorkoutDetail(arrWorkoutDetail)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllWorkout(c echo.Context) error {
	res, err := services.GetAllWorkout()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

func GetWorkout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	res, err := services.GetWorkout(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
