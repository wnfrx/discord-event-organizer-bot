package config

import (
	"errors"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	session *discordgo.Session
	db      *sqlx.DB
	router  *gin.Engine

	commands []*discordgo.ApplicationCommand
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

	if c.router == nil {
		return errors.New("router is not initialized yet")
	}

	if os.Getenv("PORT") == "" {
		if err = c.router.Run(); err != nil {
			return err
		}
	} else {
		if err = c.router.Run(":" + os.Getenv("PORT")); err != nil {
			return err
		}
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
