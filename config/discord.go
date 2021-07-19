package config

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func (c *Config) InitDiscordSession() (err error) {
	log.Println("Initializing bot session...")

	s, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
		return err
	}

	c.session = s
	return nil
}
