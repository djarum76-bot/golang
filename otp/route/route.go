package route

import (
	"os"
	"otp/controller"
	"otp/model"

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

	r.POST("/send-otp", controller.SendOTP)
	r.POST("/check-otp", controller.CheckOTP)

	return e
}
