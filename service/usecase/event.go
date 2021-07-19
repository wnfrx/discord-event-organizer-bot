package usecase

import (
	"github.com/wnfrx/discord-event-organizer-bot/models"
	"github.com/wnfrx/discord-event-organizer-bot/service"
)

type eventUsecase struct {
	er service.EventRepository
}

func NewEventUsecase(
	er service.EventRepository,
) service.EventUsecase {
	return &eventUsecase{
		er: er,
	}
}

func (u *eventUsecase) GetEvents() (result []models.Event, err error) {
	return u.er.GetEvents()
}
