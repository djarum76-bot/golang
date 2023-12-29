package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"otp/db"
	"otp/helper"
	"otp/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func Register(phone string, password string) (model.Response, error) {
	var res model.Response
	var err error
	var ID int

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO users (phone, password, is_verified) VALUES ($1, $2, $3) RETURNING ID"

	passwordHash, err := helper.HashPassword(password)
	if err != nil {
		return res, err
	}

	err = conn.QueryRow(sqlStatement, phone, passwordHash, false).Scan(&ID)
	if err != nil {
		return res, err
	}

	claims := &model.JwtCustomClaims{
		ID:    ID,
		Phone: phone,
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
	res.Message = "Success Register"
	res.Data = map[string]string{
		"token": t,
	}

	return res, nil
}

func Login(phone string, password string) (model.Response, error) {
	var res model.Response
	var err error
	var pwdHash string
	var user model.User

	conn := db.CreateConn()

	sqlStatement := "SELECT * FROM users WHERE phone = $1"

	err = conn.QueryRow(sqlStatement, phone).Scan(&user.ID, &user.Phone, &pwdHash, &user.IsVerified)
	if err != nil {
		return res, err
	}

	match, err := helper.CheckPasswordHash(pwdHash, password)
	if !match {
		return res, err
	}

	claims := &model.JwtCustomClaims{
		ID:    user.ID,
		Phone: user.Phone,
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
	res.Message = "Success Login"
	res.Data = map[string]interface{}{
		"user":  user,
		"token": t,
	}

	return res, nil
}

func SendOTP(ID int, phone string) (model.Response, error) {
	var res model.Response
	var err error

	rand.Seed(time.Now().UnixNano())

	otp := rand.Intn(1000000)

	otpString := fmt.Sprintf("%06d", otp)

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   os.Getenv("ACCOUNT_SID"),
		Password:   os.Getenv("AUTH_TOKEN"),
		AccountSid: os.Getenv("ACCOUNT_SID"),
	})

	from := "whatsapp:+14155238886"
	to := "whatsapp:" + phone
	body := "Your OTP is: " + otpString

	params := &api.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(body)

	_, err = client.Api.CreateMessage(params)
	if err != nil {
		return res, err
	}

	conn := db.CreateConn()

	sqlStatement := "INSERT INTO otps (user_id, otp, created_at, expired_at) VALUES ($1, $2, $3, $4)"

	createdAtTime := time.Now()
	createdAtString := createdAtTime.Format("2006-01-02 15:04:05")

	expiredAtTime := time.Now().Add(time.Minute * 5)
	expiredAtString := expiredAtTime.Format("2006-01-02 15:04:05")

	_, err = conn.Exec(sqlStatement, ID, otp, createdAtString, expiredAtString)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success Send OTP"
	res.Data = nil

	return res, nil
}

func CheckOTP(otp int, ID int) (model.Response, error) {
	var res model.Response
	var err error
	var otpRes model.OTP

	conn := db.CreateConn()

	sqlStatement := "SELECT otp FROM otps WHERE otp = $1 AND user_id = $2 AND expired_at > $3"

	timeNowTime := time.Now()
	timeNowString := timeNowTime.Format("2006-01-02 15:04:05")

	err = conn.QueryRow(sqlStatement, otp, ID, timeNowString).Scan(&otpRes.OTP)
	if err != nil {
		return res, err
	} else {
		sqlStatement = "UPDATE users SET is_verified = $1 WHERE id = $2"

		_, err = conn.Exec(sqlStatement, true, ID)
		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Success Check OTP"
	res.Data = nil

	return res, nil
}
