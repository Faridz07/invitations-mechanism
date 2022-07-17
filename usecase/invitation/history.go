package usecase_invitation

import (
	"errors"
	"fmt"
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/model"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (u *invitationUsecase) HistoryInvitation(claims interface{}, page, size string) (history model.History, err error) {

	newClaims := claims.(jwt.MapClaims)
	if newClaims[constant.ID] == nil {
		err = errors.New(constant.ErrSomethingWhenWrong)
		return
	}

	id, err := uuid.Parse(fmt.Sprintf("%s", newClaims[constant.ID]))
	if err != nil {
		err = errors.New(constant.ErrSomethingWhenWrong)
		return
	}

	total := u.invitationRepository.GetTotalInvitations(id)

	pageNum, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(size)
	totalPage := (int(total) + pageSize - 1) / pageSize

	start := (pageNum - 1) * pageSize

	if start > int(total) {
		start = int(total)
	}

	end := start + pageSize
	if end > int(total) {
		end = int(total)
	}

	invitations := u.invitationRepository.GetHistoryInvitation(id, pageSize, start)

	history = model.History{
		Paginate: model.Paginate{
			Page:      pageNum,
			Size:      pageSize,
			TotalRow:  int(total),
			TotalPage: totalPage,
		},
		Invitation: invitations,
	}

	return
}
