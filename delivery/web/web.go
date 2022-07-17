package web

import (
	"invitations-mechanism/delivery/middleware"
	web_users "invitations-mechanism/delivery/web/users"
	repository_users "invitations-mechanism/repository/users"
	usecase_users "invitations-mechanism/usecase/users"

	web_invitation "invitations-mechanism/delivery/web/invitation"
	repository_invitation "invitations-mechanism/repository/invitation"
	usecase_invitation "invitations-mechanism/usecase/invitation"
	"net/http"

	"invitations-mechanism/delivery/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {

	userRepo := repository_users.NewUserRepository().SetDB(db)
	usersUsecase := usecase_users.NewUserUsecase(userRepo)
	usersWeb := web_users.NewUserWeb(usersUsecase)

	invitationRepo := repository_invitation.NewInvitationRepository().SetDB(db)
	invitationUsecase := usecase_invitation.NewInvitationUsecase(invitationRepo)
	invitationWeb := web_invitation.NewInvitationWeb(invitationUsecase)

	router := gin.Default()
	router.Use(middleware.Logger())

	router.GET("/ping", func(c *gin.Context) {
		helper.ResponseOK(c, "pong!")
	})

	api := router.Group("api")

	v1 := api.Group("v1")
	v1.GET("/test", func(c *gin.Context) {
		helper.ResponseOK(c, "server Up!")
	})

	admin := v1.Group("admin")
	{
		admin.POST("/register", usersWeb.Register)
		admin.POST("/login", usersWeb.Login)
	}

	invitation := admin.Group("invitation").Use(middleware.AdminValidations())
	{
		invitation.GET("/generate", invitationWeb.GenerateInvitation)
		invitation.GET("/history", invitationWeb.HistoryInvitation)
	}

	pubInvitation := admin.Group("invitation")
	{
		pubInvitation.GET("/validate/:code", invitationWeb.ValidateInvitation)
	}

	router.NoRoute(func(c *gin.Context) {
		helper.ResponseErrorWithCode(c, http.StatusMethodNotAllowed, "method not allowed!", nil)
	})

	return router
}
