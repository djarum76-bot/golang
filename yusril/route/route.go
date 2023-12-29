package route

import (
	"os"
	"yusril/controllers"
	"yusril/models"

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

	r.POST("/person-with-image", controllers.AddPersonWithImage)
	r.POST("/person-without-image", controllers.AddPersonWithoutImage)
	r.GET("/persons", controllers.GetAllPerson)
	r.GET("/person/:id", controllers.GetPerson)
	r.GET("/person", controllers.SearchPerson)
	r.GET("/category", controllers.GetAllCategory)
	r.GET("/tags", controllers.GetAllTags)
	r.PUT("/person/:id", controllers.UpdatePerson)
	r.PUT("/image/:id", controllers.UpdateImage)
	r.PUT("/images/:id", controllers.UpdateImages)
	r.DELETE("/person/:id", controllers.DeletePerson)

	return e
}
