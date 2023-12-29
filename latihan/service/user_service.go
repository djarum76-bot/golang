package service

import (
	"latihan/db"
	"latihan/helper"
	"latihan/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func Register(email string, password string) (model.AuthResponse, error) {
	var res model.AuthResponse
	var err error
	var ID int

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	err = conn.QueryRow(sqlStatement, email, password).Scan(&ID)
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
	res.Message = "Success Register"
	res.Token = t

	return res, nil
}

func Login(email string, password string) (model.AuthResponse, error) {
	var res model.AuthResponse
	var err error
	var pwdHash string
	var user model.User

	conn := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE email = ($1)"

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
	res.Message = "Success Login"
	res.Token = t

	return res, nil
}

func GetAllUser(ID int) (model.Response, error) {
	var res model.Response
	var err error
	var user model.User
	arrUser := []model.User{}

	conn := db.CreateConn()

	sqlStatement := "SELECT id, email FROM users WHERE id != $1"

	rows, err := conn.Query(sqlStatement, ID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return res, err
		}

		arrUser = append(arrUser, user)
	}

	res.Status = http.StatusOK
	res.Message = "Success Get All Person"
	res.Data = arrUser

	return res, nil
}

func GetUser(ID int) (model.Response, error) {
	var res model.Response
	var err error
	var user model.User

	conn := db.CreateConn()

	sqlStatement := "SELECT id, email FROM users WHERE id = ($1)"

	err = conn.QueryRow(sqlStatement, ID).Scan(&user.ID, &user.Email)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Get Person"
	res.Data = user

	return res, nil
}

func DeleteUser(ID int) (model.Response, error) {
	var res model.Response
	var err error

	conn := db.CreateConn()

	sqlStatement := "DELETE FROM users WHERE id = ($1)"

	_, err = conn.Exec(sqlStatement, ID)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Delete Person"
	res.Data = nil

	return res, nil
}

func UpdateUser(ID int, email string) (model.Response, error) {
	var res model.Response
	var err error

	conn := db.CreateConn()

	sqlStatement := "UPDATE users SET email = $1 WHERE id = $2"

	_, err = conn.Exec(sqlStatement, email, ID)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Update User"
	res.Data = nil

	return res, nil
}
