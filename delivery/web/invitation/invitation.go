package invitation

import (
	usecase_invitation "invitations-mechanism/usecase/invitation"

	"github.com/gin-gonic/gin"
)

type InvitationWeb interface {
	GenerateInvitation(c *gin.Context)
	ValidateInvitation(c *gin.Context)
	HistoryInvitation(c *gin.Context)
}

type invitationWeb struct {
	uc_invitation usecase_invitation.InvitationUsecase
}

func NewInvitationWeb(uc_invitation usecase_invitation.InvitationUsecase) *invitationWeb {
	return &invitationWeb{
		uc_invitation: uc_invitation,
	}
}
