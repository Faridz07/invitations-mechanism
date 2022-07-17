package usecase_user

import (
	"invitations-mechanism/model"
	repository_invitation "invitations-mechanism/repository/invitation"
	repository_user "invitations-mechanism/repository/users"
)

type UserUsecase interface {
	Register(request model.Register, role string) error
	GeneratePassword(password string) (hashedPassword string, err error)
	Login(request model.Login, role string) (jwtResult model.ResultClaims, err error)
	LoginWithInvitationCode(code, deviceId string) (message string, err error)
	ComparePassword(hashedPassword string, password string) bool
}

type userUsecase struct {
	userRepository       repository_user.UserRepository
	invitationRepository repository_invitation.InvitationRepository
}

func NewUserUsecase(userRepo repository_user.UserRepository, invitationRepo repository_invitation.InvitationRepository) *userUsecase {
	return &userUsecase{
		userRepository:       userRepo,
		invitationRepository: invitationRepo,
	}
}
