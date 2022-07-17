package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResultClaims struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}
