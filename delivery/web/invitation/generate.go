package invitation

import (
	"errors"
	"invitations-mechanism/delivery/helper"
	"invitations-mechanism/infrastructure/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *invitationWeb) GenerateInvitation(c *gin.Context) {

	claims, exist := c.Get(constant.JWTClaims)
	if !exist {
		err := errors.New(constant.ErrMapClaimsNotFound)
		helper.ResponseErrorWithCode(c, http.StatusUnauthorized, constant.UNAUTHORIZED, err)
		return
	}

	code, err := u.uc_invitation.GenerateInvitation(claims)
	if err != nil {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrGenerateInvitation, err)
		return
	}

	helper.ResponseOKWithSingleData(c, constant.SUCCESS, code)

}
