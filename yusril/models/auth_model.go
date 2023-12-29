package models

import "github.com/golang-jwt/jwt"

type AuthResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}

type JwtCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
