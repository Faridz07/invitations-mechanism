package usecase_user

import (
	"errors"
	"fmt"
	"invitations-mechanism/infrastructure/constant"
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

	jwtResult, err = jwt.GenerateJwtToken(user)
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

func (u *userUsecase) LoginWithInvitationCode(code, deviceId string) (message string, err error) {

	retry := constant.MAX_RETRY_LOGIN_WITH_INVITATIONS
	retry, ttl := u.userRepository.GetLoginAttempt(deviceId, false, retry)
	if u.invitationRepository.CheckActiveInvitationCode(code) && retry > 0 {
		message = constant.SUCCESS_INVITATION_LOGIN
		return
	}

	retry, ttl = u.userRepository.GetLoginAttempt(deviceId, true, retry)
	if retry <= 0 && ttl.Minutes() > 0 {
		err = errors.New(fmt.Sprintf(constant.ErrToManyAttempts, ttl.String()))
		return
	} else {
		err = errors.New(fmt.Sprintf(constant.ErrRetryInvalidCode, retry))
	}

	return

}
