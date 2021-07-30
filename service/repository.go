package service

import (
	"context"

	"github.com/wnfrx/discord-event-organizer-bot/models"
)

type EventRepository interface {
	GetEvents() (result []models.Event, err error)
}

type GuildRepository interface {
	GetGuilds(ctx context.Context) (result []models.Guild, err error)
	GetActiveGuilds(ctx context.Context) (result []models.Guild, err error)
	GetGuildByID(ctx context.Context, id string) (result models.Guild, err error)
	InsertGuild(ctx context.Context, form models.FormInsertGuild) (id int64, err error)
	UpdateGuild(ctx context.Context, id string, form models.FormUpdateGuild) (err error)
}
