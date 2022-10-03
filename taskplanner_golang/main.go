package main

import (
	"taskplanner_golang/db"
	"taskplanner_golang/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":9000"))
}
