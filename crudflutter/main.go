package main

import (
	"crudflutter/db"
	"crudflutter/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":9000"))
}
