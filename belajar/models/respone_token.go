package models

type ResponseToken struct {
	Status int         `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"data"`
	Token  string      `json:"token"`
}
