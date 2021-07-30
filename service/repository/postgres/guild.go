package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/wnfrx/discord-event-organizer-bot/models"
	"github.com/wnfrx/discord-event-organizer-bot/service"
)

type guildRepository struct {
	db *sqlx.DB
}

func NewGuildRepository(
	db *sqlx.DB,
) service.GuildRepository {
	return &guildRepository{
		db: db,
	}
}

func (r *guildRepository) GetGuilds(ctx context.Context) (result []models.Guild, err error) {
	return result, nil
}

func (r *guildRepository) GetActiveGuilds(ctx context.Context) (result []models.Guild, err error) {
	return result, nil
}

func (r *guildRepository) GetGuildByID(ctx context.Context, id string) (result models.Guild, err error) {
	query := `
		SELECT
			id,
			guild_id,
			is_active,
			created_at,
			created_by,
			updated_at,
			updated_by,
			deleted_at,
			deleted_by
		FROM
			guild.guilds
		WHERE
			guild_id = $1
			AND deleted_at IS NULL
	`

	rows, err := r.db.QueryxContext(ctx, query, id)
	if err != nil && err != sql.ErrNoRows {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var scanner models.GuildScanner

		if err = rows.StructScan(&scanner); err != nil {
			return result, err
		}

		result = models.Guild{
			ID:       scanner.GuildID,
			IsActive: scanner.IsActive,
		}
	}

	return result, nil
}

func (r *guildRepository) InsertGuild(ctx context.Context, form models.FormInsertGuild) (id int64, err error) {
	query := `
		INSERT INTO
			guild.guilds
		(
			guild_id,
			is_active,
			created_at,
			created_by,
			updated_at
		)
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5
		)
		RETURNING
			id
	`

	err = r.db.QueryRowxContext(
		ctx,
		query,
		form.GuildID,
		form.IsActive,
		time.Now(),
		form.CreatedBy,
		time.Now(),
	).Scan(&id)

	if err != nil {
		return id, err
	}

	return id, err
}

func (r *guildRepository) UpdateGuild(ctx context.Context, id string, form models.FormUpdateGuild) (err error) {
	query := `
		UPDATE
			guild.guilds
		SET
			is_active = $2,
			updated_at = $3,
			updated_by = $4
		WHERE
			guild_id = $1
	`

	_, err = r.db.ExecContext(
		ctx,
		query,
		id,
		form.IsActive,
		time.Now(),
		form.UpdatedBy,
	)
	if err != nil {
		return err
	}

	return nil
}
