package users

import (
	"invitations-mechanism/infrastructure/logger"
	usecase "invitations-mechanism/usecase/users"

	"github.com/gin-gonic/gin"
)

type UserWeb interface {
	Register(c *gin.Context)
}

type userWeb struct {
	uc_users usecase.UserUsecase
}

func NewUserWeb(uc_user usecase.UserUsecase) *userWeb {
	return &userWeb{
		uc_users: uc_user,
	}
}

func (u *userWeb) Register(c *gin.Context) {
	logger.LogInfo(u.Register, "Register WEB")

	u.uc_users.Register()
}
