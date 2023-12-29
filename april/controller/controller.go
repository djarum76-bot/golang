package controller

import (
	"april/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func getTokenInfo(c echo.Context) *model.JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)

	return claims
}
