package usecase_user

import (
	"invitations-mechanism/model"
	repository_user "invitations-mechanism/repository/users"

	"github.com/google/uuid"
)

type UserUsecase interface {
	Register(request model.Register)
}

type userUsecase struct {
	userRepository repository_user.UserRepository
}

func NewUserUsecase(userRepo repository_user.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (u *userUsecase) Register(request model.Register) {

	user := model.User{
		Id:       uuid.New(),
		Username: request.Username,
		Email:    request.Email,
		Password: request.ConfirmPassword,
	}

	u.userRepository.InsertUser(user)
}
