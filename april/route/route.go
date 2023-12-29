package route

import (
	"april/controller"
	"april/model"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)

	r := e.Group("/auth")
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("SIGNING_KEY")),
	}
	r.Use(middleware.JWTWithConfig(config))

	r.GET("/user", controller.GetUser)

	r.POST("/post", controller.CreatePost)
	r.GET("/post", controller.GetAllPost)
	r.GET("/post/:ID", controller.GetPost)
	r.PUT("/post/:ID", controller.UpdatePost)
	r.DELETE("/post/:ID", controller.DeletePost)

	return e
}
