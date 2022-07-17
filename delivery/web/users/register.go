package users

import (
	"invitations-mechanism/constant"
	"invitations-mechanism/delivery/helper"
	"invitations-mechanism/model"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func (u *userWeb) Register(c *gin.Context) {
	var register model.Register
	if err := c.ShouldBind(&register); err != nil {
		if register.Username == "" || register.Email == "" || register.Password == "" || register.ConfirmPassword == "" {
			helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.InvalidRequest, err)
			return
		}
	} else if register.Password != register.ConfirmPassword {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.PasswordDoesntMatch, nil)
		return
	} else if re := regexp.MustCompile(constant.EmailRegex); !re.MatchString(register.Email) {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.InvalidEmail, nil)
		return
	}

	u.uc_users.Register(register)
}
