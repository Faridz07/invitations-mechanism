package usecase_invitation

import (
	"errors"
	"fmt"
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/model"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (u *invitationUsecase) GenerateInvitation(claims interface{}) (invitation model.ValidateInvitation, err error) {

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var length = rand.Intn(constant.MAX_LEN_INVITATION-constant.MIN_LEN_INVITATION) + constant.MIN_LEN_INVITATION
	b := make([]byte, length)
	for i := range b {
		b[i] = constant.Charset[seededRand.Intn(len(constant.Charset))]
	}

	newClaims := claims.(jwt.MapClaims)
	if newClaims[constant.ID] == nil {
		err = errors.New(constant.ErrSomethingWhenWrong)
		return
	}

	data := model.Invitation{
		Id:        uuid.New(),
		Code:      string(b),
		Status:    constant.STATUS_INVITATION_ACTIVE,
		CreatedBy: uuid.MustParse(fmt.Sprintf("%v", newClaims[constant.ID])),
		CreatedAt: time.Now().Format(time.RFC3339),
		ExpiredAt: time.Now().AddDate(0, 0, constant.EXPIRED_INVITATION).Format(time.RFC3339),
	}

	err = u.invitationRepository.InsertInvitation(data)
	if err != nil {
		return
	}

	invitation = model.ValidateInvitation{
		Code:      data.Code,
		Status:    data.Status,
		ExpiredAt: data.ExpiredAt,
	}
	return
}
