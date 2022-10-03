package main

import (
	"github.com/djarum76-bot/cukuran/db"
	"github.com/djarum76-bot/cukuran/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
