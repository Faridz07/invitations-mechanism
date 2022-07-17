package usecase_user

import (
	"invitations-mechanism/model"
	repository_user "invitations-mechanism/repository/users"
)

type UserUsecase interface {
	Register(request model.Register, role string) error
	GeneratePassword(password string) (hashedPassword string, err error)
	Login(request model.Login) (jwtResult model.ResultClaims, err error)
	ComparePassword(hashedPassword string, password string) bool
}

type userUsecase struct {
	userRepository repository_user.UserRepository
}

func NewUserUsecase(userRepo repository_user.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}
