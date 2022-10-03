package models

import (
	"database/sql"
	"net/http"
	"taskplanner_golang/db"
)

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

func GetUser(id string) (Response, error) {
	var res Response
	var user User

	con := db.CreateCon()

	sqlStatement := "SELECT id,username,image,role FROM users WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Username, &user.Image, &user.Role)
	if err == sql.ErrNoRows {
		return res, err
	}
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Get Data User"
	res.Data = user

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
