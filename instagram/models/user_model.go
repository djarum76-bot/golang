package models

import (
	"instagram/db"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func CompleteProfile(id int, phone string, picture *multipart.FileHeader) (Response, error) {
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "UPDATE users SET phone = ($1), picture = ($2) WHERE id = ($3)"

	//picture
	src, err := picture.Open()
	if err != nil {
		return res, err
	}
	defer src.Close()

	pictureUrl := "profile/" + picture.Filename

	dst, err := os.Create(pictureUrl)
	if err != nil {
		return res, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return res, err
	}

	_, err = con.Exec(sqlStatement, phone, pictureUrl, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Completing Profile Success"

	return res, nil
}

func GetUser(id int) (User, error) {
	var user User
	var err error

	con := db.CreateCon()

	sqlStatement := "SELECT id, username, email, phone, picture, createdAt FROM users WHERE id = ($1)"

	err = con.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.Picture, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
