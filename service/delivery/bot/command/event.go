package command

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/wnfrx/discord-event-organizer-bot/constants"
)

func (h *botCommandHandler) commandHandlerPing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}

func (h *botCommandHandler) commandHandlerShowEvents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   constants.DiscordMessageFlagLoading,
			Content: "Getting events...",
		},
	})

	events, err := h.euc.GetEvents()
	if err != nil {
		s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
			Content: "Something happened while getting events",
		})
		return
	}

	followUpMessage := "Hi, " + i.Member.User.Mention() + "\n"
	followUpMessage += "Yayy! Here's the event list you're requested below"

	messageFormat := `
		> ID: %d
		> UserGuildID: %d
		> ChannelID: %d
		> Name: %s
		> Description: %s
		> EventTime: %s
		> Duration: %d
	`

	for _, event := range events {
		message := fmt.Sprintf(
			messageFormat,
			event.ID,
			event.UserGuildID,
			event.ChannelID,
			event.Name,
			event.Description,
			event.EventTime,
			event.Duration,
		)

		followUpMessage += message
	}

	log.Println(followUpMessage)

	s.InteractionResponseEdit(s.State.User.ID, i.Interaction, &discordgo.WebhookEdit{
		Content: followUpMessage,
	})
}
