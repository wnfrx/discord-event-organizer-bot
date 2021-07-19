package constants

// Reference:
// https://discord.com/developers/docs/resources/channel
const (
	DiscordMessageFlagCrossposted          = 1 << 0
	DiscordMessageFlagIsCrosspost          = 1 << 1
	DiscordMessageFlagSuppressEmbeds       = 1 << 2
	DiscordMessageFlagSourceMessageDeleted = 1 << 3
	DiscordMessageFlagUrgent               = 1 << 4
	DiscordMessageFlagHasThread            = 1 << 5
	DiscordMessageFlagEphemeral            = 1 << 6
	DiscordMessageFlagLoading              = 1 << 7
)
