package models

import (
	"instagram/db"
	"net/http"
)

func CreateComment(postId int, userId int, comment string, createdAt string) (Response, error) {
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "INSERT INTO comments (post_id, user_id, comment, createdAt) VALUES ($1, $2, $3, $4)"

	_, err = con.Exec(sqlStatement, postId, userId, comment, createdAt)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Add Comment"

	return res, nil
}

func GetAllComment(postId string) ([]Comment, error) {
	arrComment := []Comment{}
	var comment Comment
	var user User

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM comments WHERE post_id = ($1) ORDER BY createdAt DESC"

	rows, err := con.Query(sqlStatement, postId)
	if err != nil {
		return arrComment, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&comment.Id, &comment.PostId, &comment.UserId, &comment.Komen, &comment.CreatedAt)

		if err != nil {
			return arrComment, err
		}

		user, err = GetUser(comment.UserId)
		if err != nil {
			return arrComment, err
		}
		comment.User = user

		arrComment = append(arrComment, comment)
	}

	return arrComment, nil
}
