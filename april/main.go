package main

import (
	"april/db"
	"april/route"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db.Init()

	e := route.Init()

	port := ":" + os.Getenv("PORT")

	e.Logger.Fatal(e.Start(port))
}