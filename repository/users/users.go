package repository_user

import (
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user model.User)
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

func (r *userRepository) InsertUser(user model.User) {
	logger.LogInfo(r.InsertUser, "Register Repository")
}
