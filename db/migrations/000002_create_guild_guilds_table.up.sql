BEGIN;

CREATE TABLE IF NOT EXISTS guild.guilds
(
    id         BIGSERIAL NOT NULL,
    guild_id   VARCHAR(100),
    is_active  BOOLEAN      DEFAULT TRUE,
    created_at TIMESTAMP    DEFAULT NOW(),
    created_by VARCHAR(100) DEFAULT 'SYSTEM'::CHARACTER VARYING,
    updated_at TIMESTAMP    DEFAULT NOW(),
    updated_by VARCHAR(100) DEFAULT 'SYSTEM'::CHARACTER VARYING,
    deleted_at TIMESTAMP,
    deleted_by VARCHAR(100) DEFAULT NULL
);

CREATE INDEX guild_guilds_guild_id_idx ON guild.guilds(guild_id);

COMMIT;
