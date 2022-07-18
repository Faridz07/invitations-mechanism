package usecase_invitation

import (
	"errors"
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"
	"time"

	"github.com/google/uuid"
)

func (u *invitationUsecase) ValidateInvitation(code string) (validation model.ValidateInvitation, err error) {
	invitation := u.invitationRepository.GetInvitationByCode(code)

	if invitation.Id == uuid.Nil {
		err = errors.New(constant.ErrInvalidCode)
		logger.LogError(u.ValidateInvitation, err.Error(), err)
		return
	}

	if invitation.Status == constant.STATUS_INVITATION_ACTIVE && invitation.ExpiredAt.Before(time.Now()) {
		results, err := u.invitationRepository.UpdateInvitation(code, constant.STATUS_INVITATION_ACTIVE, constant.STATUS_INVITATION_INACTIVE)
		if results.Id != uuid.Nil && err == nil {
			invitation = results
		}
	}

	validation = model.ValidateInvitation{
		Code:      invitation.Code,
		Status:    invitation.Status,
		ExpiredAt: invitation.ExpiredAt,
	}

	return
}
