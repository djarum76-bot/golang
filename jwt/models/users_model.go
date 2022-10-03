package models

import (
	"database/sql"
	"fmt"
	"jwt/db"
	"jwt/helpers"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username string, password string) (bool, error) {
	var obj User
	var pwdHash string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwdHash)
	if err == sql.ErrNoRows {
		return false, err
	}
	if err != nil {
		return false, err
	}

	match, err := helpers.CheckPasswordHash(pwdHash, password)
	if !match {
		fmt.Println("Hash dan password tidak sesuai")
		return false, err
	}

	return true, nil
}
