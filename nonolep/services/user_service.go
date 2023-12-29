package services

import (
	"io"
	"mime/multipart"
	"net/http"
	"nonolep/db"
	"nonolep/helpers"
	"nonolep/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lib/pq"
)

func Register(email string, password string) (models.Response, error) {
	var res models.Response
	var err error
	var ID int

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id"

	pwdHash, err := helpers.HashPassword(password)
	if err != nil {
		return res, err
	}

	err = conn.QueryRow(sqlStatement, email, pwdHash).Scan(&ID)
	if err != nil {
		return res, err
	}

	claims := &models.JwtCustomClaims{
		ID:    ID,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Register Successfull"
	res.Data = map[string]string{
		"token": t,
	}

	return res, nil
}

func Login(email string, password string) (models.Response, error) {
	var res models.Response
	var err error
	var user models.User
	var pwdHash string

	conn := db.CreateConn()

	sqlStatement := "SELECT id, email, password FROM users WHERE email = $1"

	err = conn.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Email, &pwdHash)
	if err != nil {
		return res, err
	}

	match, err := helpers.CheckPasswordHash(pwdHash, password)
	if !match {
		return res, err
	}

	claims := &models.JwtCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 720).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Login Successfull"
	res.Data = map[string]string{
		"token": t,
	}

	return res, nil
}

func GetUser(ID int) (models.Response, error) {
	var res models.Response
	var err error
	var user models.User

	conn := db.CreateConn()

	sqlStatement := "SELECT id, email, gender, age, weight, height, goals, level, picture, name, phone, is_filled FROM users WHERE id = $1"

	err = conn.QueryRow(sqlStatement, ID).Scan(&user.ID, &user.Email, &user.Gender, &user.Age, &user.Weight, &user.Height, pq.Array(&user.Goals), &user.Level, &user.Picture, &user.Name, &user.Phone, &user.IsFilled)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Get User Successfull"
	res.Data = user

	return res, nil
}

func UpdateUser(user models.User, picture *multipart.FileHeader, isEditProfile bool) (models.Response, error) {
	var res models.Response
	var err error
	var pictureURL string
	var sqlStatement string

	conn := db.CreateConn()

	if isEditProfile {
		sqlStatement = "UPDATE users SET gender = $1, age = $2, weight = $3, height = $4, goals = $5, level = $6, name = $7, phone = $8, is_filled = $9 WHERE id = $10"
	} else {
		sqlStatement = "UPDATE users SET gender = $1, age = $2, weight = $3, height = $4, goals = $5, level = $6, picture = $7, name = $8, phone = $9, is_filled = $10 WHERE id = $11"
	}

	if picture == nil {
		pictureURL = "image/default-user.png"
	} else {
		src, err := picture.Open()
		if err != nil {
			return res, err
		}
		defer src.Close()

		pictureURL = "image/" + picture.Filename

		dst, err := os.Create(pictureURL)
		if err != nil {
			return res, err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return res, err
		}
	}

	if isEditProfile {
		_, err = conn.Exec(sqlStatement, user.Gender, user.Age, user.Weight, user.Height, pq.Array(user.Goals), user.Level, user.Name, user.Phone, true, user.ID)
		if err != nil {
			return res, err
		}
	} else {
		_, err = conn.Exec(sqlStatement, user.Gender, user.Age, user.Weight, user.Height, pq.Array(user.Goals), user.Level, pictureURL, user.Name, user.Phone, true, user.ID)
		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Update Profile Successfull"
	res.Data = nil

	return res, nil
}
