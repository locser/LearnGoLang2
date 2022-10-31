package middleware

import (
	"LearnGoLang2/model"
	sercurity2 "LearnGoLang2/security"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: sercurity2.SECRET_KEY,
		Claims:     &model.JwtCustomClaims{},
	}

	return middleware.JWTWithConfig(config)
}
