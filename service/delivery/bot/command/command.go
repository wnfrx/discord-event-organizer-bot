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

	// Voting Region
	CommandVote       = "vote"
	CommandVoteCreate = "create"
	CommandVoteCancel = "cancel"
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
		{
			Name:        CommandVote,
			Description: "Vote commands",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        CommandVoteCreate,
					Description: "Buat voting gan",
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "description",
							Description: "Mau voting tentang apa gan?",
							Required:    true,
						},
						{
							Type:        discordgo.ApplicationCommandOptionInteger,
							Name:        "duration",
							Description: "Berapa detik?",
							Required:    true,
						},
					},
				},
			},
		},
	}
)
