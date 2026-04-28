package infrastructure

import (
	"go-simple-api/internal/config"
	"go-simple-api/internal/delivery/http/dto"
	"go-simple-api/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, cfg *config.Config) (*dto.Token, error) {

	secret := cfg.Jwt.JwtSecret

	claims := &JWTCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(utils.StringToInt(cfg.Jwt.JwtExpired)) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(secret))

	return &dto.Token{
		Token:     token,
		ExpiredAt: claims.ExpiresAt.Time,
	}, err
}
