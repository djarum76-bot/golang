package service

import (
	"april/db"
	"april/helper"
	"april/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func Register(name string, email string, password string) (model.Response, string, error) {
	var res model.Response
	var err error
	var ID int

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"

	pwdHash, err := helper.HashPassword(password)
	if err != nil {
		return res, err.Error(), err
	}

	err = conn.QueryRow(sqlStatement, name, email, pwdHash).Scan(&ID)
	if err != nil {
		return res, err.Error(), err
	}

	claims := &model.JwtCustomClaims{
		ID:    ID,
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return res, err.Error(), err
	}

	res.Status = http.StatusOK
	res.Message = "Register Successfull"
	res.Data = map[string]string{
		"token": t,
	}

	return res, "", nil
}

func Login(email string, password string) (model.Response, string, error) {
	var res model.Response
	var err error
	var pwdHash string
	var user model.User

	conn := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE email = $1"

	err = conn.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &pwdHash, &user.Name)
	if err != nil {
		return res, "Email atau Password salah", err
	}

	match, err := helper.CheckPasswordHash(pwdHash, password)
	if !match {
		return res, "Email atau Password salah", err
	}

	claims := &model.JwtCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return res, err.Error(), err
	}

	res.Status = http.StatusOK
	res.Message = "Login Successfull"
	res.Data = map[string]string{
		"token": t,
	}

	return res, "", nil
}

func GetUser(ID int) (model.Response, error) {
	var res model.Response
	var err error
	var user model.User

	conn := db.CreateConn()

	sqlStatement := "SELECT id, email, name FROM users WHERE id = $1"

	err = conn.QueryRow(sqlStatement, ID).Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Get User Successfull"
	res.Data = user

	return res, nil
}
