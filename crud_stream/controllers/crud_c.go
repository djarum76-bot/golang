package controllers

import (
	"crud_stream/db"
	"crud_stream/models"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	// "golang.org/x/net/websocket"
)

var upgrader = websocket.Upgrader{}

func AddUser(c echo.Context) error {
	username := c.FormValue("username")

	result, err := models.AddUser(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func RealTime(c echo.Context) error {
	var user models.User
	var arrUser []models.User = []models.User{}
	var res models.Response

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	con := db.CreateCon()

	for {
		sqlStatement := "SELECT * FROM users"

		rows, err := con.Query(sqlStatement)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&user.Id, &user.Username)
			if err != nil {
				return err
			}

			arrUser = append(arrUser, user)
		}

		res.Status = http.StatusOK
		res.Pesan = "Sukses"
		res.Data = arrUser

		err = ws.WriteJSON(res)
		if err != nil {
			c.Logger().Error(err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}

		fmt.Printf("%s\n", msg)
	}
}

// func RealTimeNet(c echo.Context) error {
// 	websocket.Handler(func(ws *websocket.Conn) {
// 		defer ws.Close()
// 		for {
// 			// Write
// 			err := websocket.Message.Send(ws, "Hello, Client!")
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}

// 			// Read
// 			msg := ""
// 			err = websocket.Message.Receive(ws, &msg)
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}
// 			fmt.Printf("%s\n", msg)
// 		}
// 	}).ServeHTTP(c.Response(), c.Request())
// 	return nil
// }
