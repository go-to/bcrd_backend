package repository

import (
	"time"

	"github.com/go-to/bcrd_backend/model"
)

type IEventRepository interface {
	GetActiveEvents(time *time.Time) (*model.ActiveEvent, error)
}

type EventRepository struct {
	model model.IEventModel
}

func NewEventRepository(m model.EventModel) *EventRepository {
	return &EventRepository{&m}
}

func (r *EventRepository) GetActiveEvents(time *time.Time) (*model.ActiveEvent, error) {
	return r.model.FindActiveEvent(time)
}
