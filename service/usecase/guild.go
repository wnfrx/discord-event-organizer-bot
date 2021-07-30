package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/wnfrx/discord-event-organizer-bot/models"
	"github.com/wnfrx/discord-event-organizer-bot/service"
)

var (
	ErrGuildNotFound = errors.New("guild not found")
)

type guildUsecase struct {
	gr service.GuildRepository
}

func NewGuildUsecase(
	gr service.GuildRepository,
) service.GuildUsecase {
	return &guildUsecase{
		gr: gr,
	}
}

func (u *guildUsecase) RegisterGuild(ctx context.Context, id string) (err error) {
	guild, err := u.gr.GetGuildByID(ctx, id)
	if err != nil {
		return err
	}

	if guild.ID == "" {
		_, err = u.gr.InsertGuild(ctx, models.FormInsertGuild{
			GuildID:   id,
			IsActive:  true,
			CreatedBy: "SYSTEM",
		})
		if err != nil {
			log.Printf("[usecase][RegisterGuild] Failed to InsertGuild, %+v\n", err)
			return err
		}
	} else {
		if !guild.IsActive {
			err = u.gr.UpdateGuild(ctx, guild.ID, models.FormUpdateGuild{
				IsActive:  true,
				UpdatedBy: "SYSTEM",
			})
			if err != nil {
				log.Printf("[usecase][RegisterGuild] Failed to UpdateGuild, %+v\n", err)
				return err
			}
		}
	}

	return nil
}

func (u *guildUsecase) RemoveGuild(ctx context.Context, id string) (err error) {
	guild, err := u.gr.GetGuildByID(ctx, id)
	if err != nil {
		return err
	}

	if guild.ID == "" {
		return ErrGuildNotFound
	}

	if guild.IsActive {
		err = u.gr.UpdateGuild(ctx, guild.ID, models.FormUpdateGuild{
			IsActive:  false,
			UpdatedBy: "SYSTEM",
		})
		if err != nil {
			log.Printf("[usecase][RemoveGuild] Failed to UpdateGuild, %+v\n", err)
			return err
		}
	}

	return nil
}
