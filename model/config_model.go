package model

import (
	"time"

	"github.com/go-to/bcrd_backend/util"
)

type Config struct {
	ID        int64
	ConfName  string
	ConfValue string
}

var config Config

func (Config) TableName() string {
	return "config"
}

type IConfigModel interface {
	GetTime() (time.Time, error)
	IsCheckEventPeriod() (bool, error)
}

type ConfigModel struct {
	db DB
}

func NewConfigModel(db DB) *ConfigModel {
	return &ConfigModel{db: db}
}

func (m *ConfigModel) GetTime() (time.Time, error) {
	confName := "debug_time"
	res := m.db.Conn.
		Model(&Config{}).
		Select("conf_value").
		Where("conf_name = ?", confName).
		Scan(&config)
	if res.Error != nil {
		return time.Time{}, res.Error
	}

	if config.ConfValue != "" {
		t, err := util.ParseTime("2006-01-02 15:04:05", config.ConfValue)
		if err != nil {
			return time.Time{}, err
		}
		return util.DateTime(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()), nil
	}

	return util.Now(), nil
}

func (m *ConfigModel) IsCheckEventPeriod() (bool, error) {
	confName := "is_check_event_period"
	res := m.db.Conn.
		Model(&Config{}).
		Select("conf_value").
		Where("conf_name = ?", confName).
		Scan(&config)
	if res.Error != nil {
		return false, res.Error
	}

	if config.ConfValue == "1" {
		return true, nil
	}
	return false, nil
}
