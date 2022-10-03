package models

import (
	"instagram/db"
	"net/http"
)

func LikePost(postId int, userId int) (Response, error) {
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "INSERT INTO likes (post_id, user_id) VALUES ($1,$2)"

	_, err = con.Exec(sqlStatement, postId, userId)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Like Post"

	return res, nil
}

func UnlikePost(id string) (Response, error) {
	var res Response
	var err error

	con := db.CreateCon()

	sqlStatement := "DELETE FROM likes WHERE id = ($1)"

	_, err = con.Exec(sqlStatement, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Unlike Post"

	return res, nil
}
