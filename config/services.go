package config

import (
	"log"

	"github.com/wnfrx/discord-event-organizer-bot/service/delivery/bot/command"
	"github.com/wnfrx/discord-event-organizer-bot/service/delivery/cron"
	"github.com/wnfrx/discord-event-organizer-bot/service/repository/faker"
	"github.com/wnfrx/discord-event-organizer-bot/service/repository/postgres"
	"github.com/wnfrx/discord-event-organizer-bot/service/usecase"
)

func (c *Config) InitServices() (err error) {
	// TODO: replace with database repository
	erf := faker.NewEventRepository()
	// er := postgres.NewEventRepository(c.db)
	gr := postgres.NewGuildRepository(c.db)

	euc := usecase.NewEventUsecase(erf)
	guc := usecase.NewGuildUsecase(gr)

	cjh := cron.NewJobHandler(c.session)

	if err := cjh.InitJobHandlers(); err != nil {
		log.Printf("[config][services] Failed while register cron job handlers, %+v", err)
		return err
	}

	bch := command.NewBotCommandHandler(
		c.session,
		euc,
		guc,
	)

	if err := bch.RegisterBotCommandHandlers(); err != nil {
		log.Printf("[config][services] Failed while register bot command handlers, %+v", err)
		return err
	}

	return nil
}
