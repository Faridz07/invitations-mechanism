package usecase_invitation

import (
	"invitations-mechanism/model"
	repository_invitation "invitations-mechanism/repository/invitation"
)

type InvitationUsecase interface {
	GenerateInvitation(claims interface{}) (invitation model.ValidateInvitation, err error)
	ValidateInvitation(code string) (validation model.ValidateInvitation, err error)
	HistoryInvitation(claims interface{}, page, size string) (invitations model.History, err error)
}

type invitationUsecase struct {
	invitationRepository repository_invitation.InvitationRepository
}

func NewInvitationUsecase(invitationRepository repository_invitation.InvitationRepository) *invitationUsecase {
	return &invitationUsecase{
		invitationRepository: invitationRepository,
	}
}
