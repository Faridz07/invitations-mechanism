package model

import (
	uuid "github.com/google/uuid"
)

type Invitation struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	Status    string    `json:"status" db:"status"`
	CreatedBy uuid.UUID `json:"created_by" db:"created_by"`
	CreatedAt string    `json:"created_at" db:"created_at"`
	ExpiredAt string    `json:"expired_at" db:"expired_at"`
}

type History struct {
	Paginate   Paginate     `json:"paginations"`
	Invitation []Invitation `json:"data"`
}

type Paginate struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalRow  int `json:"total_record"`
	TotalPage int `json:"total_page"`
}

type ValidateInvitation struct {
	Code      string `json:"code"`
	Status    string `json:"status"`
	ExpiredAt string `json:"expired_at"`
}
