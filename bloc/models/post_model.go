package models

import (
	"bloc/db"
	"net/http"
)

func AddPost(user_id int, title string, body string) (Response, error) {
	var post Post
	var res Response
	var err error
	var id int

	con := db.CreateCon()

	sqlStatement := "INSERT INTO posts (user_id, title, body) VALUES ($1, $2, $3) RETURNING id"

	err = con.QueryRow(sqlStatement, user_id, title, body).Scan(&id)
	if err != nil {
		return res, err
	}

	post.Id = id
	post.UserId = user_id
	post.Title = title
	post.Body = body

	res.Status = http.StatusOK
	res.Data = post

	return res, nil
}

func GetAllPost() (Response, error) {
	var post Post
	arrPost := []Post{}
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM posts ORDER BY id ASC"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
		if err != nil {
			return res, err
		}

		arrPost = append(arrPost, post)
	}

	res.Status = http.StatusOK
	res.Data = arrPost

	return res, nil
}

func GetPost(id string) (Response, error) {
	var res Response
	var err error
	var post Post

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM posts WHERE id = ($1)"

	err = con.QueryRow(sqlStatement, id).Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Data = post

	return res, nil
}

func DeletePost(id string) (Response, error) {
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "DELETE FROM posts WHERE id = ($1)"

	_, err = con.Exec(sqlStatement, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Data = "Sukses"

	return res, nil
}

func UpdatePost(id int, title string, body string) (Response, error) {
	var res Response
	var err error
	var post Post

	con := db.CreateCon()

	sqlStatement := "UPDATE posts SET title = ($1), body = ($2) WHERE id = ($3)"

	_, err = con.Exec(sqlStatement, title, body, id)
	if err != nil {
		return res, err
	}

	post.Id = id
	post.Title = title
	post.Body = body

	res.Status = http.StatusOK
	res.Data = post

	return res, nil
}
