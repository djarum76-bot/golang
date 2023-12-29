package service

import (
	"mei/db"
	"mei/helper"
	"mei/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func Register(email string, password string) (model.Response, error) {
	var res model.Response
	var err error
	var ID int

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	pwdHash, err := helper.HashPassword(password)
	if err != nil {
		return res, err
	}

	err = conn.QueryRow(sqlStatement, email, pwdHash).Scan(&ID)
	if err != nil {
		return res, err
	}

	claims := &model.JwtCustomClaims{
		ID:    ID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Register Successfull"
	res.Data = map[string]string{
		"token": t,
	}

	return res, nil
}

func Login(email string, password string) (model.Response, error) {
	var res model.Response
	var err error
	var user model.User
	var pwdHash string

	conn := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE email = $1"

	err = conn.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &pwdHash)
	if err != nil {
		return res, err
	}

	match, err := helper.CheckPasswordHash(pwdHash, password)
	if !match {
		return res, err
	}

	claims := &model.JwtCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Login Successfull"
	res.Data = map[string]string{
		"token": t,
	}

	return res, nil
}
