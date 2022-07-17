package users

import (
	usecase "invitations-mechanism/usecase/users"

	"github.com/gin-gonic/gin"
)

type UserWeb interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	LoginWithInvitationCode(c *gin.Context)
}

type userWeb struct {
	uc_users usecase.UserUsecase
}

func NewUserWeb(uc_user usecase.UserUsecase) *userWeb {
	return &userWeb{
		uc_users: uc_user,
	}
}
