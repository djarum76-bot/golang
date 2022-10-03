package models

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"wppb/db"
)

type Wisata struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	OpenDay     string `json:"openDay"`
	OpenTime    string `json:"openTime"`
	TicketPrice string `json:"ticketPrice"`
	Image       string `json:"image"`
	IsDone      string `json:"isDone"`
}

func AddWisata(name string, location string, description string, openDay string, openTime string, ticketPrice string, image *multipart.FileHeader) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO wisata values (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		fmt.Println("1")
		return res, err
	}

	src, err := image.Open()
	if err != nil {
		fmt.Println("2")
		return res, err
	}
	defer src.Close()

	imgUrl := "upload/" + image.Filename

	dst, err := os.Create(imgUrl)
	if err != nil {
		fmt.Println("3")
		return res, err
	}
	defer dst.Close()

	result, err := stmt.Exec(nil, name, location, description, openDay, openTime, ticketPrice, imgUrl, "0")
	if err != nil {
		fmt.Println("4")
		return res, err
	}

	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("5")
		return res, err
	}

	getIDLast, err := result.LastInsertId()
	if err != nil {
		fmt.Println("6")
		return res, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Sukses"
	res.Data = map[string]int64{
		"ID Wisata": getIDLast,
	}

	return res, nil
}

func GetAllWisata() (Response, error) {
	var wisata Wisata
	var arrWisata []Wisata = []Wisata{}
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM wisata"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&wisata.Id, &wisata.Name, &wisata.Location, &wisata.Description, &wisata.OpenDay, &wisata.OpenTime, &wisata.TicketPrice, &wisata.Image, &wisata.IsDone)
		if err != nil {
			return res, err
		}

		arrWisata = append(arrWisata, wisata)
	}

	res.Status = http.StatusOK
	res.Pesan = "Sukses"
	res.Data = arrWisata

	return res, nil
}
