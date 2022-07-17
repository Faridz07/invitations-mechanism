package users

import (
	"invitations-mechanism/delivery/helper"
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/model"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func (u *userWeb) Login(c *gin.Context) {
	var login model.Login
	if err := c.ShouldBind(&login); err != nil {
		if login.Email == "" || login.Password == "" {
			helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrInvalidRequest, err)
			return
		}
	} else if re := regexp.MustCompile(constant.EmailRegex); !re.MatchString(login.Email) {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrInvalidEmail, nil)
		return
	}

	token, err := u.uc_users.Login(login, constant.ADMIN)
	if err != nil {
		if strings.Contains(err.Error(), constant.UNAUTHORIZED) {
			helper.ResponseErrorWithCode(c, http.StatusUnauthorized, err.Error(), nil)
			return
		}

		helper.ResponseErrorWithCode(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseOKWithSingleData(c, constant.SUCCESS, token)
}

func (u *userWeb) LoginWithInvitationCode(c *gin.Context) {
	code, exist := c.Params.Get("code")
	if code == "" || !exist {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrCodeEmpty, nil)
		return
	}

	deviceId := c.Request.Header.Get("deviceId")
	if deviceId == "" {
		helper.ResponseErrorWithCode(c, http.StatusMethodNotAllowed, constant.ErrInvalidDeviceId, nil)
		return
	}

	message, err := u.uc_users.LoginWithInvitationCode(code, deviceId)
	if err != nil {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	helper.ResponseOKWithSingleData(c, constant.SUCCESS, message)
}
