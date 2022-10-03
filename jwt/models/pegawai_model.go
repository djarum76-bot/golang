package models

import (
	"jwt/db"
	"net/http"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func GetAllPegawai() (Response, error) {
	var pegawai Pegawai
	var arrPegawai []Pegawai = []Pegawai{}
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Alamat, &pegawai.Telepon)
		if err != nil {
			return res, err
		}

		arrPegawai = append(arrPegawai, pegawai)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrPegawai

	return res, nil
}

func GetPegawai(id string) (Response, error) {
	var pegawai Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(&pegawai.Id, &pegawai.Nama, &pegawai.Alamat, &pegawai.Telepon)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = pegawai

	return res, nil
}

func AddPegawai(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO pegawai values(?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nil, nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]string{"message": "Data berhasil ditambahkan"}

	return res, nil
}

func UpdatePegawai(id string, nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE pegawai set nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]string{"message": "Data berhasil diubah"}

	return res, nil
}

func DeletePegawai(id string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM pegawai WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]string{
		"message": "Data berhasil dihapus",
	}

	return res, nil
}
