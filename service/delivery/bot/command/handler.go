package command

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/wnfrx/discord-event-organizer-bot/service"
)

type botCommandHandler struct {
	session *discordgo.Session
	euc     service.EventUsecase
}

func NewBotCommandHandler(
	session *discordgo.Session,
	euc service.EventUsecase,
) botCommandHandler {
	return botCommandHandler{
		session: session,
		euc:     euc,
	}
}

func (h botCommandHandler) RegisterBotCommandHandlers() (err error) {
	var iCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		CommandPing:      h.commandHandlerPing,
		CommandShowEvent: h.commandHandlerShowEvents,
	}

	// NOTE: Interaction Command Handler
	h.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := iCommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	for _, v := range commands {
		cmd, err := h.session.ApplicationCommandCreate(os.Getenv("BOT_APPLICATION_ID"), os.Getenv("GUILD_ID"), v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v\n", v.Name, err)
			return err
		}

		log.Printf("Command %s successfully registered, ID:%s\n", cmd.Name, cmd.ID)
	}

	// NOTE: Add here for Other Command Handlers

	return nil
}
