package models

import (
	"database/sql"
	"time"
)

type (
	GuildScanner struct {
		ID        int64          `db:"id"`
		GuildID   string         `db:"guild_id"`
		IsActive  bool           `db:"is_active"`
		CreatedAt time.Time      `db:"created_at"`
		CreatedBy string         `db:"created_by"`
		UpdatedAt time.Time      `db:"updated_at"`
		UpdatedBy string         `db:"updated_by"`
		DeletedAt *time.Time     `db:"deleted_at"`
		DeletedBy sql.NullString `db:"deleted_by"`
	}

	Guild struct {
		ID              string
		Name            string
		SystemChannelID string
		IsActive        bool
	}

	FormInsertGuild struct {
		GuildID   string
		IsActive  bool
		CreatedBy string
	}

	FormUpdateGuild struct {
		IsActive  bool
		UpdatedBy string
	}
)
