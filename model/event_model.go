package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID        int64
	Name      string
	Year      int32
	StartDate time.Time
	EndDate   time.Time
}

type ActiveEvent Event

var activeEvent ActiveEvent

func (Event) TableName() string {
	return "events"
}

type IEventModel interface {
	FindActiveEvent(time *time.Time) (*ActiveEvent, error)
}

type EventModel struct {
	db DB
}

func NewEventModel(db DB) *EventModel {
	return &EventModel{db: db}
}

func (m *EventModel) FindActiveEvent(time *time.Time) (*ActiveEvent, error) {
	t := time.Format("2006-01-02")

	query := m.db.Conn.
		Model(&Event{}).
		Debug().
		Where("start_date <= ?", t).
		Where("end_date >= ?", t)

	activeEvent = ActiveEvent{}
	res := query.First(&activeEvent)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &ActiveEvent{}, nil
		}
		return nil, res.Error
	}
	return &activeEvent, nil
}
