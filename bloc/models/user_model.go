package models

import (
	"bloc/db"
	"bloc/helper"
	"database/sql"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func Register(email string, password string) (ResponseToken, error) {
	var res ResponseToken
	var err error
	var user User
	var id int

	con := db.CreateCon()

	sqlStatement := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	err = con.QueryRow(sqlStatement, email, password).Scan(&id)
	if err != nil {
		return res, err
	}

	user.Id = id
	user.Email = email

	claims := &JwtCustomClaims{
		user.Id,
		user.Email,
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
	res.Data = user
	res.Token = t

	return res, nil
}

func Login(email string, password string) (ResponseToken, error) {
	var user User
	var res ResponseToken
	var err error
	var pwdHash string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE email = ($1)"

	err = con.QueryRow(sqlStatement, email).Scan(&user.Id, &user.Email, &pwdHash)
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

	claims := &JwtCustomClaims{
		user.Id,
		user.Email,
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
	res.Data = user
	res.Token = t

	return res, nil
}

func GetUser(id int) (Response, error) {
	var user User
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "SELECT id, email FROM users WHERE id = ($1)"

	err = con.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Email)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Data = user

	return res, nil
}
