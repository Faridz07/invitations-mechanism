package repository_user

import (
	"errors"
	"invitations-mechanism/constant"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user model.User) error
	CheckUser(user model.User) (exist bool, err error)
	GetUserByEmail(email string) (result model.User, err error)
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

func (r *userRepository) InsertUser(user model.User) error {
	exist, err := r.CheckUser(user)
	if err != nil {
		return err
	}

	if exist {
		msg := errors.New("error insert data to db")
		err := errors.New("user already exist")
		logger.LogError(r.InsertUser, msg.Error(), err)
		return err
	}

	if err := r.db.Create(&user).Error; err != nil {
		msg := errors.New("error when insert data to db")
		logger.LogError(r.InsertUser, msg.Error(), err)
		return msg
	}

	return nil
}

func (r *userRepository) CheckUser(user model.User) (exist bool, err error) {
	users := []model.User{}

	if err = r.db.Where("username = ? OR email = ?", strings.ToLower(user.Username), strings.ToLower(user.Email)).Find(&users).Error; err != nil {
		msg := errors.New("error when select data from db")
		logger.LogError(r.InsertUser, msg.Error(), err)
		err = msg
		return
	}

	if len(users) > 0 {
		exist = true
		return
	}

	return
}

func (r *userRepository) GetUserByEmail(email string) (result model.User, err error) {
	if err = r.db.Where("email = ?", strings.ToLower(email)).First(&result).Error; err != nil {
		msg := errors.New("error when select data from db")
		logger.LogError(r.InsertUser, msg.Error(), err)
		err = msg
		return
	}

	if result.Id == uuid.Nil {
		err = errors.New(constant.UserDoesntExist)
		logger.LogError(r.InsertUser, err.Error(), err)
		return
	}

	return
}
