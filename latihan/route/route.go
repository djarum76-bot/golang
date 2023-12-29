package route

import (
	"latihan/controller"
	"latihan/model"
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

	r.GET("/users", controller.GetAllUser)
	r.GET("/user", controller.GetUserProfile)
	r.GET("/user/:id", controller.GetUser)
	r.DELETE("/user", controller.DeleteUser)
	r.PUT("/user", controller.UpdateUser)

	return e
}
