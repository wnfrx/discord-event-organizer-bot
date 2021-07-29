package config

import (
	"log"

	"github.com/wnfrx/discord-event-organizer-bot/service/delivery/bot/command"
	"github.com/wnfrx/discord-event-organizer-bot/service/delivery/cron"
	"github.com/wnfrx/discord-event-organizer-bot/service/repository/faker"
	"github.com/wnfrx/discord-event-organizer-bot/service/usecase"
)

func (c *Config) InitServices() (err error) {
	// TODO: replace with real repository
	er := faker.NewEventRepository()
	euc := usecase.NewEventUsecase(er)

	bch := command.NewBotCommandHandler(
		c.session,
		euc,
	)

	if err := bch.RegisterBotCommandHandlers(); err != nil {
		log.Printf("[config][services] Failed while register bot command handlers, %+v", err)
		return err
	}

	cjh := cron.NewJobHandler(c.session)

	if err := cjh.InitJobHandlers(); err != nil {
		log.Printf("[config][services] Failed while register cron job handlers, %+v", err)
		return err
	}

	return nil
}
