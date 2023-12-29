package model

import "github.com/golang-jwt/jwt"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JwtCustomClaims struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

type User struct {
	ID         int    `json:"id"`
	Phone      string `json:"phone"`
	IsVerified bool   `json:"is_verified"`
}

type OTP struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	OTP       int    `json:"otp"`
	CreatedAt string `json:"created_at"`
	ExpiredAt string `json:"expired_at"`
}
