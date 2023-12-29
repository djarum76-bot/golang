package services

import (
	"database/sql"
	"net/http"
	"os"
	"time"
	"yusril/db"
	"yusril/helper"
	"yusril/models"

	"github.com/golang-jwt/jwt"
)

func Register(email string, password string) (models.AuthResponse, error) {
	var res models.AuthResponse
	var err error
	var ID int

	conn := db.CreateCon()

	sqlStatement := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	err = conn.QueryRow(sqlStatement, email, password).Scan(&ID)
	if err != nil {
		return res, err
	}

	emailDecrypt, err := helper.Decrypt(email, os.Getenv("SECRET_STRING"))
	if err != nil {
		return res, err
	}

	claims := &models.JwtCustomClaims{
		ID:    ID,
		Email: emailDecrypt,
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
	res.Message = "Success Register"
	res.Data = nil
	res.Token = t

	return res, nil
}

func Login(email string, password string) (models.AuthResponse, error) {
	var res models.AuthResponse
	var err error
	var pwdHash string
	var user models.User

	conn := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE email = ($1)"

	err = conn.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &pwdHash)
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

	emailDecrypt, err := helper.Decrypt(email, os.Getenv("SECRET_STRING"))
	if err != nil {
		return res, err
	}

	claims := &models.JwtCustomClaims{
		ID:    user.ID,
		Email: emailDecrypt,
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
	res.Message = "Success Login"
	res.Data = nil
	res.Token = t

	return res, nil
}
