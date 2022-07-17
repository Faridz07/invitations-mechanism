package repository_user

import (
	"errors"
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user model.User) error
	CheckUser(user model.User) (exist bool, err error)
	GetUserByEmail(email string) (result model.User, err error)
	GetLoginAttempt(deviceId string, cache bool, retry int) (count int, ttl time.Duration)
}

type userRepository struct {
	db  *gorm.DB
	rdc *redis.Client
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) SetDB(db *gorm.DB) *userRepository {
	r.db = db
	return r
}

func (r *userRepository) SetRedis(rdc *redis.Client) *userRepository {
	r.rdc = rdc
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
		err = errors.New(constant.ErrUserDoesntExist)
		logger.LogError(r.InsertUser, err.Error(), err)
		return
	}

	return
}

func (r *userRepository) GetLoginAttempt(deviceId string, cache bool, retry int) (count int, ttl time.Duration) {

	count, _ = r.rdc.Get(deviceId).Int()
	ttl, _ = r.rdc.TTL(deviceId).Result()
	if count != 0 {
		retry = count
	} else if ttl.Seconds() < 0 {
		count = retry
	}

	if cache {
		if count < 0 {
			return
		} else if count == 0 && ttl.Minutes() < 0 {
			count = constant.MAX_RETRY_LOGIN_WITH_INVITATIONS - 1
			err := r.rdc.Set(deviceId, count, time.Duration(1800*time.Second)).Err()
			if err != nil {
				logger.LogError(r.GetLoginAttempt, "failed set to redis", err)
				return
			}
		} else {
			count -= 1
			err := r.rdc.Del(deviceId).Err()
			if err == nil {
				err = r.rdc.Set(deviceId, count, time.Duration(1800*time.Second)).Err()
				if err != nil {
					logger.LogError(r.GetLoginAttempt, "failed set to redis", err)
					return
				}
			}
		}
	}

	return
}
