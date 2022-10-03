package main

import (
	"instagram/db"
	"instagram/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
