package dto

import "time"

type Token struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type TokenRequest struct {
	ApiKey string `json:"api_key"`
}

type TokenResponse struct {
	Token     string    `json:"token"`
	Type      string    `json:"type"`
	ExpiredAt time.Time `json:"expired_at"`
}
