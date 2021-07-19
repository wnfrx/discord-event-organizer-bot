package service

import "github.com/wnfrx/discord-event-organizer-bot/models"

type EventRepository interface {
	GetEvents() (result []models.Event, err error)
}
