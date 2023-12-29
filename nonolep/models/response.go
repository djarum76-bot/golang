package models

import "github.com/golang-jwt/jwt"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JwtCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	ID       int      `json:"id"`
	Email    string   `json:"email"`
	Gender   string   `json:"gender"`
	Age      int      `json:"age"`
	Weight   int      `json:"weight"`
	Height   int      `json:"height"`
	Goals    []string `json:"goals"`
	Level    string   `json:"level"`
	Picture  string   `json:"picture"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
	IsFilled bool     `json:"is_filled"`
}

type Workout struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Duration      int             `json:"duration"`
	Level         string          `json:"level"`
	Picture       string          `json:"picture"`
	WorkoutDetail []WorkoutDetail `json:"workout_detail"`
}

type WorkoutDetail struct {
	ID        int    `json:"id"`
	WorkoutID int    `json:"workout_id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Picture   string `json:"picture"`
}
