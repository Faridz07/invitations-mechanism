package web

import (
	"invitations-mechanism/delivery/middleware"
	users_web "invitations-mechanism/delivery/web/users"
	repository_users "invitations-mechanism/repository/users"
	usecase_users "invitations-mechanism/usecase/users"
	"net/http"

	"invitations-mechanism/delivery/helper"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *gin.Engine {

	userRepo := repository_users.NewUserRepository().SetDB(db)
	usersUsecase := usecase_users.NewUserUsecase(userRepo)
	usersWeb := users_web.NewUserWeb(usersUsecase)

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
		invitation.POST("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "TEST")
		})
		invitation.GET("/check")
		invitation.GET("/history")
	}

	router.NoRoute(func(c *gin.Context) {
		helper.ResponseErrorWithCode(c, http.StatusMethodNotAllowed, "method not allowed!", nil)
	})

	return router
}
