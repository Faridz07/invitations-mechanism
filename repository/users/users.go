package repository_user

import (
	"invitations-mechanism/infrastructure/logger"

	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser()
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) SetDB(db *gorm.DB) *userRepository {
	r.db = db
	return r
}

func (r *userRepository) InsertUser() {
	logger.LogInfo(r.InsertUser, "Register Repository")
}
