package config

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	session  *discordgo.Session
	commands []*discordgo.ApplicationCommand
	db       *sqlx.DB
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Run() (err error) {
	if c.session == nil {
		return errors.New("session is not initialized yet")
	}

	err = c.session.Open()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Stop() (err error) {
	err = c.session.Close()
	if err != nil {
		return err
	}

	return nil
}
