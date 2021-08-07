package command

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/wnfrx/discord-event-organizer-bot/models"
	"github.com/wnfrx/discord-event-organizer-bot/service"
)

type botCommandHandler struct {
	guilds map[string]*discordgo.Guild

	session *discordgo.Session
	euc     service.EventUsecase
	guc     service.GuildUsecase

	// Vote
	votingGuildMap  map[string]models.Voting
	votingCancelMap map[string]*chan int
}

func NewBotCommandHandler(
	session *discordgo.Session,
	euc service.EventUsecase,
	guc service.GuildUsecase,
) *botCommandHandler {
	return &botCommandHandler{
		guilds:          map[string]*discordgo.Guild{},
		session:         session,
		euc:             euc,
		guc:             guc,
		votingGuildMap:  make(map[string]models.Voting),
		votingCancelMap: make(map[string]*chan int),
	}
}

func (h *botCommandHandler) RegisterBotCommandHandlers() (err error) {
	var iCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		CommandPing:  h.commandHandlerPing,
		CommandEvent: h.commandHandlerEvent,
		CommandVote:  h.commandHandlerVoting,
	}

	// NOTE: Interaction Command Handler
	h.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := iCommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// NOTE: Handler on Bot Online
	h.session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		for _, g := range r.Guilds {
			h.guilds[g.ID] = g

			err := h.guc.RegisterGuild(context.Background(), g.ID)
			if err != nil {
				log.Printf("[service][delivery][bot][command] Failed to RegisterGuild, %+v\n", err)
				return
			}

			for _, v := range commands {
				cmd, err := h.session.ApplicationCommandCreate(os.Getenv("BOT_APPLICATION_ID"), g.ID, v)
				if err != nil {
					log.Printf("Cannot create '%v' command on Guild [%s]: %v\n", v.Name, g.ID, err)
					return
				}

				log.Printf("Command %s successfully registered on Guild [%s], ID:%s\n", cmd.Name, g.ID, cmd.ID)
			}
		}
	})

	// NOTE: Handler on Bot joined a Guild
	h.session.AddHandler(func(s *discordgo.Session, g *discordgo.GuildCreate) {
		if _, ok := h.guilds[g.ID]; ok {
			return
		}

		h.guilds[g.ID] = g.Guild

		err := h.guc.RegisterGuild(context.Background(), g.ID)
		if err != nil {
			log.Printf("[service][delivery][bot][command] Failed to RegisterGuild, %+v\n", err)
			return
		}

		fmt.Printf("Me joined Guild #%s :D\n", g.ID)

		for _, v := range commands {
			cmd, err := h.session.ApplicationCommandCreate(os.Getenv("BOT_APPLICATION_ID"), g.ID, v)
			if err != nil {
				log.Printf("Cannot create '%v' command on Guild [%s]: %v\n", v.Name, g.ID, err)
				return
			}

			log.Printf("Command %s successfully registered on Guild [%s], ID:%s\n", cmd.Name, g.ID, cmd.ID)
		}

		// s.ChannelMessageSend(g.SystemChannelID, "Hello World!")
	})

	// NOTE: Handler on Bot kicked from a Guild
	h.session.AddHandler(func(s *discordgo.Session, g *discordgo.GuildDelete) {
		if _, ok := h.guilds[g.ID]; !ok {
			return
		}

		err := h.guc.RemoveGuild(context.Background(), g.ID)
		if err != nil {
			log.Printf("[service][delivery][bot][command] Failed to RemoveGuild, %+v\n", err)
			return
		}

		fmt.Printf("Me is kicked from Guild #%s :(\n", g.ID)

		delete(h.guilds, g.ID)
	})

	// NOTE: Add here for Other Command Handlers

	return nil
}
