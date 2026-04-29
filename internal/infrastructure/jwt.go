package infrastructure

import (
	"go-simple-api/internal/delivery/http/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(secret string, expireHours int) (*dto.Token, error) {

	// nanti ini diganti pakai request body / validasi DB
	userID := uint(1)
	// secret := cfg.Jwt.JwtSecret

	claims := &JWTCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHours) * time.Hour)),
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
