package usecase_user

import (
	"errors"
	"invitations-mechanism/constant"
	"invitations-mechanism/infrastructure/jwt"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"

	"golang.org/x/crypto/bcrypt"
)

func (u *userUsecase) Login(request model.Login, role string) (jwtResult model.ResultClaims, err error) {
	user, err := u.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		return
	}

	if !u.ComparePassword(user.Password, request.Password) {
		err = errors.New(constant.ErrLoginFailed)
		logger.LogError(u.Login, constant.ErrLoginFailed, errors.New(constant.ErrHashPasswordDoesntMatch))
		return
	}

	if user.RoleName != role {
		err = errors.New(constant.ErrUserLoginUnAuthorized)
		logger.LogError(u.Login, constant.ErrUserLoginUnAuthorized, errors.New(constant.ErrUserLoginUnAuthorized))
		return
	}

	jwtResult, err = jwt.GenerateJwtToken(user.Username, user.Email)
	if err != nil {
		msg := errors.New(constant.ErrSomethingWhenWrong)
		logger.LogError(u.Login, msg.Error(), err)
		err = msg
		return
	}

	return
}

func (u *userUsecase) ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}
