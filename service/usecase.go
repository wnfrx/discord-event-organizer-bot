package service

import "github.com/wnfrx/discord-event-organizer-bot/models"

type EventUsecase interface {
	GetEvents() (result []models.Event, err error)
}
