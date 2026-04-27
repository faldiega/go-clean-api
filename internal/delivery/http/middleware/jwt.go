package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTMiddleware struct {
	jwtSecret     string
	whitelistURLs []string
}

func NewJWTMiddleware(secret string, whitelistURLs []string) *JWTMiddleware {
	return &JWTMiddleware{
		jwtSecret:     secret,
		whitelistURLs: whitelistURLs,
	}
}

func (m *JWTMiddleware) Middleware() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			requestPath := c.Request().URL.Path

			// cek whitelist dulu
			if IsUnauthorizedURL(requestPath, m.whitelistURLs) {
				return next(c) // skip JWT, langsung lanjut
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "missing token",
				})
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid token format",
				})
			}

			tokenString := parts[1]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Validasi algoritma
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(m.jwtSecret), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid or expired token",
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "invalid claims",
				})
			}

			// simpan ke context untuk dipakai di handler
			c.Set("user_id", claims["user_id"])

			return next(c)
		}
	}
}
