package models

import (
	"crud_stream/db"
	"net/http"
)

func AddUser(username string) (Response, error) {
	var user User
	var err error
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT into users values (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nil, username)
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
	res.Pesan = "Berhasil Tambah Note"
	res.Data = user

	return res, err
}
