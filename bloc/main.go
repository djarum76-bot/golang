package main

import (
	"bloc/db"
	"bloc/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
