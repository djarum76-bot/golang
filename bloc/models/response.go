package models

import "github.com/golang-jwt/jwt"

type ResponseToken struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Token  string      `json:"token"`
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type JwtCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
