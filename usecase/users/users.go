package usecase_user

import (
	"invitations-mechanism/infrastructure/logger"
	repository_user "invitations-mechanism/repository/users"
)

type UserUsecase interface {
	Register()
}

type userUsecase struct {
	userRepository repository_user.UserRepository
}

func NewUserUsecase(userRepo repository_user.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (u *userUsecase) Register() {
	logger.LogInfo(u.Register, "Register Usecase")

	u.userRepository.InsertUser()
}
