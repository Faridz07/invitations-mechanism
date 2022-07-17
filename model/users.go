package model

import uuid "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"Username"`
	Email    string    `json:"Email"`
	Password string    `json:"Password"`
	RoleId   uuid.UUID `json:"RoleId"`
}
