package routes

import (
	"instagram/controllers"
	"instagram/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Static("/profile", "profile")
	e.Static("/post", "post")

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	r := e.Group("/auth")
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.PUT("/complete-profile", controllers.CompleteProfile)

	r.GET("/user", controllers.GetUser)

	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.GetAllPost)
	r.GET("/post/:id", controllers.GetPost)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.POST("/like/:post_id", controllers.LikePost)
	r.DELETE("/like/:id", controllers.UnlikePost)

	r.POST("/comment/:post_id", controllers.CreateComment)
	r.GET("/comment/:post_id", controllers.GetAllComment)

	return e
}
