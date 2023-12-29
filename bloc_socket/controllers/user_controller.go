package controllers

import (
	"bloc_socket/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func GetAllUser(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			if string(msg) != "" {
				c.Logger().Error(err)
			}
		}

		if string(msg) != "" {
			var user models.User
			err = json.Unmarshal(msg, &user)
			if err != nil {
				c.Logger().Error(err)
			}

			if user.Email != "" && user.Password != "" {
				err = models.AddUser(user.Email, user.Password)
				if err != nil {
					c.Logger().Error(err)
				}
			}
		}

		arrUser, err := models.GetAllUser()
		if err != nil {
			c.Logger().Error(err)
		}

		jsonUser, err := json.Marshal(arrUser)
		if err != nil {
			c.Logger().Error(err)
		}

		err = ws.WriteMessage(websocket.TextMessage, jsonUser)
		if err != nil {
			c.Logger().Error(err)
		}

		// arrUser, err := models.GetAllUser()
		// if err != nil {
		// 	c.Logger().Error(err)
		// }

		// err = ws.WriteJSON(arrUser)
		// if err != nil {
		// 	c.Logger().Error(err)
		// }
	}
}

func AddUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	err := models.AddUser(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}

// func GetAllUser(c echo.Context) error {
// 	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer ws.Close()

// 	for {
// 		arrUser, err := models.GetAllUser()
// 		if err != nil {
// 			c.Logger().Error(err)
// 		}

// 		err = ws.WriteJSON(arrUser)
// 		if err != nil {
// 			c.Logger().Error(err)
// 		}

// 		// Read
// 		err = ws.ReadJSON(arrUser)
// 		if err != nil {
// 			c.Logger().Error(err)
// 		}
// 	}
// }
