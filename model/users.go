package model

import uuid "github.com/google/uuid"

type Users struct {
	Id       string    `json:"id"`
	Username string    `json:"Username"`
	Email    string    `json:"Email"`
	Password string    `json:"Password"`
	RoleId   uuid.UUID `json:"RoleId"`
}
