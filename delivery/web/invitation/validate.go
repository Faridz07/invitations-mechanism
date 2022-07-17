package invitation

import (
	"invitations-mechanism/delivery/helper"
	"invitations-mechanism/infrastructure/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *invitationWeb) ValidateInvitation(c *gin.Context) {
	code := c.Param("code")
	if code == "" || len(code) < constant.MIN_LEN_INVITATION || len(code) > constant.MAX_LEN_INVITATION {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrInvalidCode, nil)
		return
	}

	validation, err := u.uc_invitation.ValidateInvitation(code)
	if err != nil {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrInvalidCode, err)
		return
	}

	helper.ResponseOKWithSingleData(c, constant.SUCCESS, validation)
}
