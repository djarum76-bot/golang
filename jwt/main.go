package main

import (
	"jwt/db"
	"jwt/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":9000"))
}
