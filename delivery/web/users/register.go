package users

import (
	"invitations-mechanism/constant"
	"invitations-mechanism/delivery/helper"
	"invitations-mechanism/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *userWeb) Register(c *gin.Context) {
	var register model.Register
	if err := c.ShouldBind(&register); err != nil {
		if register.Username == "" || register.Email == "" || register.Password == "" {
			helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.InvalidRequest, err)
			return
		}
	}

	u.uc_users.Register()
}
