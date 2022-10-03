package models

import (
	"crudflutter/db"
	"net/http"
)

type Barang struct {
	Id         int    `json:"id"`
	Kodebarang string `json:"kodebarang"`
	Namabarang string `json:"namabarang"`
}

func DataBarang() (Response, error) {
	var obj Barang
	var arrobj []Barang = []Barang{}
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM tabelbarang"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Kodebarang, &obj.Namabarang)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	return res, nil
}

func DataBarangTertentu(id string) (Response, error) {
	var obj Barang
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM tabelbarang where id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(&obj.Id, &obj.Kodebarang, &obj.Namabarang)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = obj

	return res, nil
}

func SimpanBarang(kodebarang string, namabarang string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO tabelbarang values (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nil, kodebarang, namabarang)
	if err != nil {
		return res, err
	}

	getIDLast, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"ID Barang": getIDLast,
	}

	return res, nil
}

func UpdateBarang(id string, kodebarang string, namabarang string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE tabelbarang set kodebarang = ?, namabarang = ? where id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(kodebarang, namabarang, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int{
		"Status": http.StatusOK,
	}

	return res, nil
}

func DeleteBarang(id string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE from tabelbarang where id = ?"

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
	res.Data = map[string]int{
		"Status": http.StatusOK,
	}

	return res, nil
}
