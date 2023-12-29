package models

import "bloc_socket/db"

func GetAllUser() ([]User, error) {
	var user User
	arrUser := []User{}
	var err error

	con := db.CreateConn()

	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return arrUser, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return arrUser, err
		}

		arrUser = append(arrUser, user)
	}

	return arrUser, nil
}

func AddUser(email string, password string) error {
	var err error

	con := db.CreateConn()

	sqlStatement := "INSERT INTO users (email, password) VALUES ($1, $2)"

	_, err = con.Exec(sqlStatement, email, password)
	if err != nil {
		return err
	}

	return nil
}
