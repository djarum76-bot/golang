package models

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Response struct {
	Status int         `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"data"`
}
