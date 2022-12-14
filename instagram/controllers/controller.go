package controllers

import (
	"instagram/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func getNilaiToken(c echo.Context) *models.JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)

	return claims
}
