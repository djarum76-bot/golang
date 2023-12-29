package model

import (
	"github.com/golang-jwt/jwt"
)

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

type Post struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Post      string `json:"post"`
	CreatedAt string `json:"created_at"`
}
