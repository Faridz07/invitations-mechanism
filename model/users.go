package model

import uuid "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"password" db:"password"`
	RoleName string    `json:"role_name" db:"role_name"`
}
