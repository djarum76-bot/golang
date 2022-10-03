package models

import (
	"database/sql"
	"instagram/db"
	"instagram/helper"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func Register(username string, email string, password string, createdAt string) (ResponseToken, error) {
	var res ResponseToken
	var err error
	var id int

	con := db.CreateCon()

	sqlStatement := "INSERT INTO users (username, email, password, createdAt) VALUES ($1, $2, $3, $4) RETURNING id"

	err = con.QueryRow(sqlStatement, username, email, password, createdAt).Scan(&id)
	if err != nil {
		return res, err
	}

	claims := &JwtCustomClaims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Register Success"
	res.Id = id
	res.Token = t

	return res, nil
}

func Login(email string, password string) (ResponseToken, error) {
	var user User
	var res ResponseToken
	var err error
	var pwdHash string

	con := db.CreateCon()

	sqlStatement := "SELECT id, password FROM users WHERE email = ($1)"

	err = con.QueryRow(sqlStatement, email).Scan(&user.Id, &pwdHash)
	if err == sql.ErrNoRows {
		return res, err
	}
	if err != nil {
		return res, err
	}

	match, err := helper.CheckPasswordHash(pwdHash, password)
	if !match {
		return res, err
	}

	claims := JwtCustomClaims{
		user.Id,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Login Success"
	res.Id = user.Id
	res.Token = t

	return res, nil
}
