package service

import (
	"context"

	"github.com/wnfrx/discord-event-organizer-bot/models"
)

type EventUsecase interface {
	GetEvents() (result []models.Event, err error)
}

type GuildUsecase interface {
	RegisterGuild(ctx context.Context, id string) (err error)
	RemoveGuild(ctx context.Context, id string) (err error)
}
