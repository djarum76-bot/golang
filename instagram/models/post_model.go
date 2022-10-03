package models

import (
	"instagram/db"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func CreatePost(userId int, image *multipart.FileHeader, caption string, createdAt string) (Response, error) {
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "INSERT INTO posts (user_id, image, caption, createdAt) VALUES ($1, $2, $3, $4)"

	src, err := image.Open()
	if err != nil {
		return res, err
	}
	defer src.Close()

	imageUrl := "post/" + image.Filename

	dst, err := os.Create(imageUrl)
	if err != nil {
		return res, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return res, err
	}

	_, err = con.Exec(sqlStatement, userId, imageUrl, caption, createdAt)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Creating Post Success"

	return res, nil
}

func GetAllPost() ([]PostUser, error) {
	arrPostUser := []PostUser{}
	var postUser PostUser
	arrLikes := []Like{}
	var like Like
	var post Post
	var user User
	var err error

	con := db.CreateCon()

	sqlStatement1 := "SELECT * FROM posts ORDER BY createdAt DESC"
	sqlStatement2 := "SELECT * FROM likes WHERE post_id = ($1)"

	rows, err := con.Query(sqlStatement1)
	if err != nil {
		return arrPostUser, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&post.Id, &post.UserId, &post.Image, &post.Caption, &post.CreatedAt)
		if err != nil {
			return arrPostUser, err
		}

		user, err = GetUser(post.UserId)
		if err != nil {
			return arrPostUser, err
		}

		rowsLike, err := con.Query(sqlStatement2, post.Id)
		if err != nil {
			return arrPostUser, err
		}
		defer rowsLike.Close()

		for rowsLike.Next() {
			err = rowsLike.Scan(&like.Id, &like.PostId, &like.UserId)

			if err != nil {
				return arrPostUser, err
			}

			arrLikes = append(arrLikes, like)
		}

		postUser.Post = post
		postUser.Post.Likes = arrLikes
		postUser.User = user

		arrLikes = []Like{}

		arrPostUser = append(arrPostUser, postUser)
	}

	return arrPostUser, nil
}

func GetPost(id string) (PostUser, error) {
	var postUser PostUser
	var post Post
	var user User
	var err error
	arrLikes := []Like{}
	var like Like

	con := db.CreateCon()

	sqlStatement1 := "SELECT * FROM posts WHERE id = ($1)"
	sqlStatement2 := "SELECT * FROM likes WHERE post_id = ($1)"

	err = con.QueryRow(sqlStatement1, id).Scan(&post.Id, &post.UserId, &post.Image, &post.Caption, &post.CreatedAt)
	if err != nil {
		return postUser, err
	}

	rows, err := con.Query(sqlStatement2, post.Id)
	if err != nil {
		return postUser, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&like.Id, &like.PostId, &like.UserId)

		if err != nil {
			return postUser, err
		}

		arrLikes = append(arrLikes, like)
	}

	user, err = GetUser(post.UserId)
	if err != nil {
		return postUser, err
	}

	postUser.Post = post
	postUser.Post.Likes = arrLikes
	postUser.User = user

	return postUser, nil
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
	res.Message = "Deleting Success"

	return res, nil
}
