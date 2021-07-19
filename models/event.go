package models

type Event struct {
	ID          int64  `db:"id"`
	UserGuildID int64  `db:"user_guild_id"`
	ChannelID   int64  `db:"channel_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	EventTime   string `db:"event_time"`
	Duration    int64  `db:"duration"`
}
