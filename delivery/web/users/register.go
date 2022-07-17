package users

import (
	"invitations-mechanism/delivery/helper"
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/model"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func (u *userWeb) Register(c *gin.Context) {
	var register model.Register
	if err := c.ShouldBind(&register); err != nil {
		if register.Username == "" || register.Email == "" || register.Password == "" || register.ConfirmPassword == "" {
			helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrInvalidRequest, err)
			return
		}
	} else if register.Password != register.ConfirmPassword {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrPasswordDoesntMatch, nil)
		return
	} else if re := regexp.MustCompile(constant.EmailRegex); !re.MatchString(register.Email) {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrInvalidEmail, nil)
		return
	}

	err := u.uc_users.Register(register, constant.ADMIN)
	if err != nil {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helper.ResponseOKWithSingleData(c, constant.SUCCESS, nil)
}
