package main

import (
	"nonolep/db"
	"nonolep/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db.Init()

	e := routes.Init()

	port := ":" + os.Getenv("PORT")

	e.Logger.Fatal(e.Start(port))
}
