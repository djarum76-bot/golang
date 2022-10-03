package models

import (
	"belajar/db"
	"belajar/helper"
	"database/sql"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type JwtCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func Register(username string, password string) (Response, error) {
	var user User
	var err error
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT into users values (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nil, username, password)
	if err != nil {
		return res, err
	}

	getID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	user.Id = int(getID)
	user.Username = username

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Register"
	res.Data = user

	return res, nil
}

func Login(username string, password string) (ResponseToken, bool, error) {
	var user User
	var pwdHash string
	var res ResponseToken

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(&user.Id, &user.Username, &pwdHash)
	if err == sql.ErrNoRows {
		return res, false, err
	}
	if err != nil {
		return res, false, err
	}

	match, err := helper.CheckPasswordHash(pwdHash, password)
	if !match {
		return res, false, err
	}

	claims := &JwtCustomClaims{
		user.Id,
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return res, false, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Login"
	res.Data = user
	res.Token = t

	return res, true, nil
}

func GetAllUser() (Response, error) {
	var user User
	var arrUser []User = []User{}
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id,username FROM users"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username)
		if err != nil {
			return res, err
		}

		arrUser = append(arrUser, user)
	}

	res.Status = http.StatusOK
	res.Pesan = "Sukses"
	res.Data = arrUser

	return res, nil
}

func DeleteUser(id string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM users WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Berhasil dihapus"
	res.Data = map[string]string{
		"message": "Data berhasil dihapus",
	}

	return res, nil
}
