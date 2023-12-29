package main

import (
	gomail "gopkg.in/gomail.v2"
)

func main() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "butsjems6@gmail.com")
	msg.SetHeader("To", "ravipantek4@gmail.com")
	msg.SetHeader("Subject", "Hello")
	msg.SetBody("text/html", "<b>Hi banh</b>")

	n := gomail.NewDialer("smtp.gmail.com", 587, "butsjems6@gmail.com", "habil123")

	err := n.DialAndSend(msg)
	if err != nil {
		panic(err.Error())
	}
}
