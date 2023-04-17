package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hotel-booking-api/model"
	"os"
)

func JWTMiddleWare() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}
	return middleware.JWTWithConfig(config)
}

func JWTRefreshMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("SECRET_REFRESH_KEY")),
		Claims:     &model.JwtCustomClaims{},
	}
	return middleware.JWTWithConfig(config)
}
