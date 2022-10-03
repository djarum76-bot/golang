package routes

import (
	"belajar/controllers"
	"belajar/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
	e.POST("/upload", controllers.Upload)

	// need token
	r := e.Group("/restricted")
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/users", controllers.GetAllUser)
	r.GET("/print", controllers.CetakToken)
	r.DELETE("/users", controllers.DeleteUser)

	return e
}
