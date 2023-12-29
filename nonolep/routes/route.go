package routes

import (
	"nonolep/controllers"
	"nonolep/models"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Static("/image", "image")

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	r := e.Group("/auth")
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("SIGNING_KEY")),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/user", controllers.GetUser)
	r.PUT("/user", controllers.UpdateUser)

	r.POST("/workout", controllers.AddWorkout)
	r.POST("/workout-detail", controllers.AddWorkoutDetail)
	r.GET("/workout", controllers.GetAllWorkout)
	r.GET("/workout/:id", controllers.GetWorkout)

	return e
}
