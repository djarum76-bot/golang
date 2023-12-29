package service

import (
	"maret/db"
	"maret/helper"
	"maret/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func Register(email string, password string, name string) (model.Response, error) {
	var res model.Response
	var err error
	var ID int

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id"

	passwordHash, err := helper.HashPassword(password)
	if err != nil {
		return res, err
	}

	err = conn.QueryRow(sqlStatement, email, passwordHash, name).Scan(&ID)
	if err != nil {
		return res, err
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
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Register"
	res.Data = map[string]string{
		"token": t,
	}

	return res, nil
}

func Login(email string, password string) (model.Response, error) {
	var res model.Response
	var err error
	var pwdHash string
	var user model.User

	conn := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE email = $1"

	err = conn.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &pwdHash, &user.Name)
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
		Name:  user.Name,
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
	res.Data = map[string]interface{}{
		"user":  user,
		"token": t,
	}

	return res, nil
}

func GetAllUser() (model.Response, error) {
	var res model.Response
	var err error
	var user model.User
	arrUser := []model.User{}

	conn := db.CreateConn()

	sqlStatement := "SELECT id, email, name FROM users"

	rows, err := conn.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Name)
		if err != nil {
			return res, err
		}

		arrUser = append(arrUser, user)
	}

	res.Status = http.StatusOK
	res.Message = "Success Get All User"
	res.Data = arrUser

	return res, nil
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
	res.Message = "Success Get User"
	res.Data = user

	return res, nil
}

func UpdateUser(ID int, email string, name string) (model.Response, error) {
	var res model.Response
	var err error

	conn := db.CreateConn()

	sqlStatement := "UPDATE users SET email = $1, name = $2 WHERE id = $3"

	_, err = conn.Exec(sqlStatement, email, name, ID)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Update User"
	res.Data = nil

	return res, nil
}

func DeleteUser(ID int) (model.Response, error) {
	var res model.Response
	var err error

	conn := db.CreateConn()

	sqlStatement := "DELETE FROM users WHERE id = $1"

	_, err = conn.Exec(sqlStatement, ID)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Update User"
	res.Data = nil

	return res, nil
}
