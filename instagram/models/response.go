package models

import (
	"database/sql"

	"github.com/golang-jwt/jwt"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseToken struct {
	Status  int    `json:"status"`
	Id      int    `json:"id"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type User struct {
	Id        int            `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Phone     sql.NullString `json:"phone"`
	Picture   sql.NullString `json:"picture"`
	CreatedAt string         `json:"createdAt"`
}

type JwtCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type Post struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Image     string `json:"image"`
	Caption   string `json:"caption"`
	CreatedAt string `json:"createdAt"`
	Likes     []Like `json:"likes"`
}

type PostUser struct {
	Post Post `json:"post"`
	User User `json:"user"`
}

type Like struct {
	Id     int `json:"id"`
	PostId int `json:"post_id"`
	UserId int `json:"user_id"`
}

type Comment struct {
	Id        int    `json:"id"`
	PostId    int    `json:"post_id"`
	UserId    int    `json:"user_id"`
	Komen     string `json:"komen"`
	CreatedAt string `json:"createdAt"`
	User      User   `json:"user"`
}
