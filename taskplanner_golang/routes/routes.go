package routes

import (
	"taskplanner_golang/controllers"
	"taskplanner_golang/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// need token
	r := e.Group("/auth")
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("/user", controllers.GetUser)

	//task
	r.POST("/task", controllers.AddTask)
	r.GET("/task", controllers.GetAllTask)
	r.GET("/taskdate", controllers.GetAllTaskDate)
	r.GET("/task/:id", controllers.GetTask)
	r.DELETE("/task/:id", controllers.DeleteTask)
	r.PUT("/task/:id", controllers.UpdateTask)

	//note
	r.POST("/note", controllers.AddNote)
	r.GET("/note", controllers.GetAllNote)
	r.GET("/note/:id", controllers.GetNote)
	r.DELETE("/note/:id", controllers.DeleteNote)
	r.PUT("/note/:id", controllers.UpdateNote)

	return e
}
