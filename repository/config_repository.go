package repository

import (
	"time"

	"github.com/go-to/bcrd_backend/model"
)

type IConfigRepository interface {
	GetTime() (time.Time, error)
	IsCheckEventPeriod() (bool, error)
}

type ConfigRepository struct {
	model model.IConfigModel
}

func NewConfigRepository(m model.ConfigModel) *ConfigRepository {
	return &ConfigRepository{&m}
}

func (r *ConfigRepository) GetTime() (time.Time, error) {
	return r.model.GetTime()
}

func (r *ConfigRepository) IsCheckEventPeriod() (bool, error) {
	return r.model.IsCheckEventPeriod()
}
