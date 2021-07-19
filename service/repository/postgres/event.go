package postgres

import (
	"github.com/wnfrx/discord-event-organizer-bot/models"
	"github.com/wnfrx/discord-event-organizer-bot/service"
)

type eventRepository struct {
}

func NewEventRepository() service.EventRepository {
	return &eventRepository{}
}

func (r *eventRepository) GetEvents() (result []models.Event, err error) {

	return result, nil
}
