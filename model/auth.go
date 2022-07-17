package model

type Register struct {
	Username string `json:"Username" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type Login struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
