package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/wnfrx/discord-event-organizer-bot/models"
	"github.com/wnfrx/discord-event-organizer-bot/service"
)

type eventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(
	db *sqlx.DB,
) service.EventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) GetEvents() (result []models.Event, err error) {

	return result, nil
}
