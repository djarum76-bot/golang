package services

import (
	"io"
	"mime/multipart"
	"net/http"
	"nonolep/db"
	"nonolep/models"
	"os"
)

func AddWorkout(workout models.Workout, picture *multipart.FileHeader) (models.Response, error) {
	var res models.Response
	var err error
	var pictureURL string
	duration := workout.Duration
	levels := []string{"Beginner", "Intermediate", "Advanced"}

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO workouts (name, duration, level, picture) VALUES ($1, $2, $3, $4)"

	src, err := picture.Open()
	if err != nil {
		return res, err
	}
	defer src.Close()

	pictureURL = "image/" + picture.Filename

	dst, err := os.Create(pictureURL)
	if err != nil {
		return res, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return res, err
	}

	for _, level := range levels {
		_, err = conn.Exec(sqlStatement, workout.Name, duration, level, pictureURL)
		if err != nil {
			return res, err
		}

		duration += 3
	}

	res.Status = http.StatusOK
	res.Message = "Add Workout Successful"
	res.Data = nil

	return res, nil
}

func AddWorkoutDetail(workoutDetailList []models.WorkoutDetail) (models.Response, error) {
	var res models.Response
	var err error

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO workout_details (workout_id, name, duration, picture) VALUES ($1, $2, $3, $4)"

	for _, workoutDetail := range workoutDetailList {
		_, err = conn.Exec(sqlStatement, workoutDetail.WorkoutID, workoutDetail.Name, workoutDetail.Duration, workoutDetail.Picture)
		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Add Workout Detail Successful"
	res.Data = nil

	return res, nil
}

func GetAllWorkout() (models.Response, error) {
	var res models.Response
	var err error
	var workout models.Workout
	arrWorkout := []models.Workout{}

	conn := db.CreateConn()

	sqlStatement := "SELECT * FROM workouts"

	rows, err := conn.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&workout.ID, &workout.Name, &workout.Duration, &workout.Level, &workout.Picture)
		if err != nil {
			return res, err
		}

		workout.WorkoutDetail = []models.WorkoutDetail{}

		arrWorkout = append(arrWorkout, workout)
	}

	res.Status = http.StatusOK
	res.Message = "Get All Workout Successful"
	res.Data = arrWorkout

	return res, nil
}

func GetWorkout(id int) (models.Response, error) {
	var res models.Response
	var err error
	var workout models.Workout
	var workoutDetail models.WorkoutDetail
	arrWorkoutDetail := []models.WorkoutDetail{}

	conn := db.CreateConn()

	sqlStatement1 := "SELECT * FROM workouts WHERE id = $1"

	err = conn.QueryRow(sqlStatement1, id).Scan(&workout.ID, &workout.Name, &workout.Duration, &workout.Level, &workout.Picture)
	if err != nil {
		return res, err
	}

	sqlStatement2 := "SELECT * FROM workout_details WHERE workout_id = $1"

	rows, err := conn.Query(sqlStatement2, id)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&workoutDetail.ID, &workoutDetail.WorkoutID, &workoutDetail.Name, &workoutDetail.Duration, &workoutDetail.Picture)
		if err != nil {
			return res, err
		}

		arrWorkoutDetail = append(arrWorkoutDetail, workoutDetail)
	}

	workout.WorkoutDetail = arrWorkoutDetail

	res.Status = http.StatusOK
	res.Message = "Get Workout Successful"
	res.Data = workout

	return res, nil
}
