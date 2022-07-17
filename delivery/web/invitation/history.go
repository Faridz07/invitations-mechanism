package invitation

import (
	"errors"
	"invitations-mechanism/delivery/helper"
	"invitations-mechanism/infrastructure/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *invitationWeb) HistoryInvitation(c *gin.Context) {

	page := c.Request.URL.Query().Get("page")
	size := c.Request.URL.Query().Get("size")

	if page == "" {
		page = constant.DEFAULT_PAGE
	}

	if size == "" {
		size = constant.DEFAULT_SIZE
	}

	claims, exist := c.Get(constant.JWTClaims)
	if !exist {
		err := errors.New(constant.ErrMapClaimsNotFound)
		helper.ResponseErrorWithCode(c, http.StatusUnauthorized, constant.UNAUTHORIZED, err)
		return
	}

	history, err := u.uc_invitation.HistoryInvitation(claims, page, size)
	if err != nil {
		helper.ResponseErrorWithCode(c, http.StatusBadRequest, constant.ErrGetInvitation, err)
		return
	}

	helper.ResponseOKWithSingleData(c, constant.SUCCESS, history)
}
