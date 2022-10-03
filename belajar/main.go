package main

import (
	"belajar/db"
	"belajar/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":9000"))
}
