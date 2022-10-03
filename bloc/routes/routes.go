package routes

import (
	"bloc/controllers"
	"bloc/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	r := e.Group("/auth")
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/user", controllers.GetUser)

	r.POST("/post", controllers.AddPost)
	r.GET("/posts", controllers.GetAllPost)
	r.GET("/post/:id", controllers.GetPost)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.PUT("/post/:id", controllers.UpdatePost)

	return e
}
