package model

import (
	"github.com/golang-jwt/jwt"
)

type AuthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

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
	ID    int    `json:"id"`
	Email string `json:"email"`
}
