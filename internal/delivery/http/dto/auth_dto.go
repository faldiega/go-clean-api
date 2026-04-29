package dto

import "time"

type Token struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type TokenRequest struct {
	ApiKey string `json:"api_key" validate:"required,min=32"`
}

type TokenResponse struct {
	Token       string    `json:"token"`
	Type        string    `json:"type"`
	ExpireHours int       `json:"expire_hour"`
	ExpiredAt   time.Time `json:"expired_at"`
}
