package usecase_user

import (
	"errors"
	"invitations-mechanism/constant"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) Register(request model.Register, role string) error {

	hashPassword, err := u.GeneratePassword(request.ConfirmPassword)
	if err != nil {
		return err
	}

	if role != "" {
		role = constant.USER
	}

	user := model.User{
		Id:       uuid.New(),
		Username: strings.ToLower(request.Username),
		Email:    strings.ToLower(request.Email),
		Password: hashPassword,
		RoleName: role,
	}

	err = u.userRepository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) GeneratePassword(password string) (hashedPassword string, msg error) {

	byteHashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		msg = errors.New("failed to hash password")
		logger.LogError(u.GeneratePassword, msg.Error(), err)
		return
	}

	err = bcrypt.CompareHashAndPassword(byteHashPassword, []byte(password))

	if err != nil {
		msg = errors.New("failed to compare password")
		logger.LogError(u.GeneratePassword, msg.Error(), err)
		return
	}

	hashedPassword = string(byteHashPassword)
	return
}
