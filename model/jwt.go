package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type MyClaims struct {
	jwt.StandardClaims
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}

type ResultClaims struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}
