package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return c.JSON(401, "missing token")
			}

			split := strings.Split(authHeader, " ")
			if len(split) != 2 || split[0] != "Bearer" {
				return c.JSON(401, "invalid token format")
			}

			tokenString := split[1]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Validasi algoritma
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(401, "invalid token")
			}

			claims := token.Claims.(jwt.MapClaims)

			// inject ke context
			c.Set("user_id", claims["user_id"])

			return next(c)
		}
	}
}
