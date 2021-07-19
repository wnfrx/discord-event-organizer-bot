package command

import "github.com/bwmarrin/discordgo"

const (
	CommandPing      = "ping"
	CommandShowEvent = "show-event"
	CommandAddEvent  = "add-event"
	CommandJoinEvent = "join-event"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        CommandPing,
			Description: "Ping me!",
		},
		{
			Name:        CommandShowEvent,
			Description: "Show all events",
		},
	}
)
