package models

import (
	"database/sql"

	"github.com/golang-jwt/jwt"
)

type ResponseToken struct {
	Status int         `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"data"`
	Token  string      `json:"token"`
}

type Response struct {
	Status int         `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"data"`
}

type User struct {
	Id       int            `json:"id"`
	Username string         `json:"username"`
	Image    sql.NullString `json:"image"`
	Role     sql.NullString `json:"role"`
}

type JwtCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type Task struct {
	Id        int    `json:"id"`
	User_Id   int    `json:"user_id"`
	Title     string `json:"title"`
	Tanggal   string `json:"tanggal"`
	Waktu     string `json:"waktu"`
	Date      string `json:"date"`
	Deskripsi string `json:"deskripsi"`
}

type TaskDate struct {
	Tanggal string `json:"tanggal"`
	Task    []Task `json:"task"`
}

type Note struct {
	Id      int    `json:"id"`
	User_Id int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
