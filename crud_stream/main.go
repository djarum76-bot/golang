package main

import (
	"crud_stream/db"
	"crud_stream/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":1323"))
}
