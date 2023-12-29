package main

import (
	"bloc_socket/db"
	"bloc_socket/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
