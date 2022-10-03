package models

type Response struct {
	Status int         `json:"status"`
	Pesan  string      `json:"pesan"`
	Data   interface{} `json:"data"`
}
