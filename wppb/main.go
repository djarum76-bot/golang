package main

import (
	"wppb/db"
	"wppb/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":9000"))
}
