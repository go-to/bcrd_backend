package repository

import (
	"time"

	"github.com/go-to/bcrd_backend/model"
)

type IStampRepository interface {
	AddStamp(time *time.Time, userId string, shopId int64) (int32, error)
	DeleteStamp(userId string, shopId int64) (int32, error)
}

type StampRepository struct {
	model model.IStampModel
}

func NewStampRepository(m model.StampModel) *StampRepository {
	return &StampRepository{model: &m}
}

func (r *StampRepository) AddStamp(time *time.Time, userId string, shopId int64) (int32, error) {
	return r.model.AddStamp(time, userId, shopId)
}

func (r *StampRepository) DeleteStamp(userId string, shopId int64) (int32, error) {
	return r.model.DeleteStamp(userId, shopId)
}
