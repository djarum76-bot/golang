package service

import (
	"april/db"
	"april/model"
	"net/http"
)

func CreatePost(userID int, post string, createdAt string) error {
	var err error

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO posts (user_id, post, created_at) VALUES ($1, $2, $3)"

	_, err = conn.Exec(sqlStatement, userID, post, createdAt)
	if err != nil {
		return err
	}

	return nil
}

func GetAllPost(userID int) (model.Response, error) {
	var res model.Response
	var err error
	var post model.Post
	arrPost := []model.Post{}

	conn := db.CreateConn()

	sqlStatement := "SELECT * FROM posts WHERE user_id = $1"

	rows, err := conn.Query(sqlStatement, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&post.ID, &post.UserID, &post.Post, &post.CreatedAt)
		if err != nil {
			return res, err
		}

		arrPost = append(arrPost, post)
	}

	res.Status = http.StatusOK
	res.Message = "Get All Post Successfull"
	res.Data = arrPost

	return res, nil
}

func GetPost(userID int, ID int) (model.Response, error) {
	var res model.Response
	var err error
	var post model.Post

	conn := db.CreateConn()

	sqlStatement := "SELECT * FROM posts WHERE user_id = $1 AND id = $2"

	err = conn.QueryRow(sqlStatement, userID, ID).Scan(&post.ID, &post.UserID, &post.Post, &post.CreatedAt)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Get Post Successfull"
	res.Data = post

	return res, nil
}

func UpdatePost(user_id int, id int, post string) error {
	var err error

	conn := db.CreateConn()

	sqlStatement := "UPDATE posts SET post = $1 WHERE id = $2 AND user_id = $3"

	_, err = conn.Exec(sqlStatement, post, id, user_id)
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(user_id int, id int) error {
	var err error

	conn := db.CreateConn()

	sqlStatement := "DELETE FROM posts WHERE user_id = $1 AND id = $2"

	_, err = conn.Exec(sqlStatement, user_id, id)
	if err != nil {
		return err
	}

	return nil
}
