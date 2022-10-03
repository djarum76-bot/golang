package models

import (
	"net/http"
	"strconv"
	"taskplanner_golang/db"
)

func AddNote(userId int, title string, content string) (Response, error) {
	var note Note
	var err error
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT into notes values (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nil, userId, title, content)
	if err != nil {
		return res, err
	}

	getID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	note.Id = int(getID)
	note.User_Id = userId
	note.Title = title
	note.Content = content

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Tambah Note"
	res.Data = note

	return res, err
}

func GetAllNote(userId int) (Response, error) {
	var note Note
	var arrNote []Note = []Note{}
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM notes WHERE user_id = ?"

	rows, err := con.Query(sqlStatement, userId)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&note.Id, &note.User_Id, &note.Title, &note.Content)
		if err != nil {
			return res, err
		}

		arrNote = append(arrNote, note)
	}

	res.Status = http.StatusOK
	res.Pesan = "Sukses"
	res.Data = arrNote

	return res, nil
}

func GetNote(id string) (Response, error) {
	var note Note
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM notes WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(&note.Id, &note.User_Id, &note.Title, &note.Content)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Sukses"
	res.Data = note

	return res, nil
}

func DeleteNote(id string, userId int) (Response, error) {
	var res Response
	var note Note

	con := db.CreateCon()

	sqlStatement := "DELETE FROM notes WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return res, err
	}

	note.Id, _ = strconv.Atoi(id)
	note.User_Id = userId

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Hapus Data"
	res.Data = note

	return res, err
}

func UpdateNote(id string, userId int, title string, content string) (Response, error) {
	var note Note
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "UPDATE notes set title = ?, content = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(title, content, id)
	if err != nil {
		return res, err
	}

	note.Id, _ = strconv.Atoi(id)
	note.User_Id = userId
	note.Title = title
	note.Content = content

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Update Data"
	res.Data = note

	return res, err
}
