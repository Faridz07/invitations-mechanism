package repository_invitation

import (
	"errors"
	"invitations-mechanism/infrastructure/constant"
	"invitations-mechanism/infrastructure/logger"
	"invitations-mechanism/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvitationRepository interface {
	InsertInvitation(invitation model.Invitation) error
	UpdateInvitation(code string, status string, updatedStatus string) (result model.Invitation, err error)
	CheckInvitationCode(code string) bool
	CheckActiveInvitationCode(code string) bool
	GetInvitationByCode(code string) (invitation model.Invitation)
	GetHistoryInvitation(id uuid.UUID, limit, offset int) (invitations []model.Invitation)
	GetTotalInvitations(id uuid.UUID) (count int64)
}

type invitationRepository struct {
	db *gorm.DB
}

func NewInvitationRepository() *invitationRepository {
	return &invitationRepository{}
}

func (r *invitationRepository) SetDB(db *gorm.DB) *invitationRepository {
	r.db = db
	return r
}

func (r *invitationRepository) InsertInvitation(invitation model.Invitation) error {

	if r.CheckActiveInvitationCode(invitation.Code) {
		msg := errors.New("failed insert data to db")
		err := errors.New("code already exist")
		logger.LogError(r.InsertInvitation, msg.Error(), err)
		return err
	}

	if err := r.db.Create(&invitation).Error; err != nil {
		msg := errors.New("error when insert data to db")
		logger.LogError(r.InsertInvitation, msg.Error(), err)
		return msg
	}

	return nil
}

func (r *invitationRepository) CheckActiveInvitationCode(code string) bool {
	invitation := model.Invitation{}

	if rows := r.db.Where("code = ? and status = ? and expired_at between now()::timestamptz and now()::timestamptz + INTERVAL '7 DAY'", code, constant.STATUS_INVITATION_ACTIVE).First(&invitation).RowsAffected; rows > 0 {
		return true
	}

	if r.CheckInvitationCode(code) {
		invitation, err := r.UpdateInvitation(code, constant.STATUS_INVITATION_ACTIVE, constant.STATUS_INVITATION_INACTIVE)
		if err != nil || invitation.Status == constant.STATUS_INVITATION_ACTIVE {
			return true
		}
	}

	return false
}

func (r *invitationRepository) CheckInvitationCode(code string) bool {
	invitation := model.Invitation{}
	if rows := r.db.Where("code = ? and status = ?", code, constant.STATUS_INVITATION_ACTIVE).First(&invitation).RowsAffected; rows > 0 {
		return true
	}

	return false
}

func (r *invitationRepository) UpdateInvitation(code string, status string, updatedStatus string) (result model.Invitation, err error) {
	if err = r.db.Model(&result).Where("code = ? and status = ?", code, status).Update("status", updatedStatus).Error; err != nil {
		logger.LogError(r.UpdateInvitation, "failed when update to db", err)
		return
	}

	result = r.GetInvitationByCode(code)

	return
}

func (r *invitationRepository) GetInvitationByCode(code string) (invitation model.Invitation) {
	if rows := r.db.Where("code = ?", code).Order("created_at desc").First(&invitation).RowsAffected; rows > 0 {
		return
	}

	return
}

func (r *invitationRepository) GetTotalInvitations(id uuid.UUID) (count int64) {
	result := r.db.Select("count(*) as count").Where("created_by = ?", id).Order("created_at desc").Model(&model.Invitation{}).Count(&count)
	if err := result.Error; err != nil {
		logger.LogError(r.GetTotalInvitations, "error when counting to db", err)
		return
	}

	return
}

func (r *invitationRepository) GetHistoryInvitation(id uuid.UUID, limit, offset int) (invitations []model.Invitation) {
	if err := r.db.Order("created_at desc").Limit(limit).Offset(offset).Find(&invitations).Error; err != nil {
		logger.LogError(r.GetHistoryInvitation, "error when select to db", err)
		return
	}

	return
}
