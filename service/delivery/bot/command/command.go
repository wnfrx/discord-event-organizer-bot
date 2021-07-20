package command

import "github.com/bwmarrin/discordgo"

const (
	CommandPing = "ping"

	// Event Region
	CommandEvent        = "event"
	CommandEventGetList = "list"
	CommandEventCreate  = "create"
	CommandEventJoin    = "join"
	CommandEventCancel  = "cancel"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        CommandPing,
			Description: "Ping me!",
		},
		{
			Name:        CommandEvent,
			Description: "Event commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        CommandEventGetList,
					Description: "Show event list",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "filter",
							Description: "List Filter: upcoming, ongoing, past, all",
							Required:    false,
						},
					},
				},
			},
		},
	}
)
