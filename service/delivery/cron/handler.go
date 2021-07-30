package cron

import (
	"fmt"
	"log"
	"os"
)

func (h *cronJobHandler) testJob() {
	fmt.Println("Every minutes")

	guildID := os.Getenv("GUILD_ID")

	if guildID == "" {
		guildID = "687140071887208484"
	}

	guild, err := h.session.State.Guild(guildID)
	if err != nil {
		log.Printf("[cron][handler] Failed while get guild, %+v\n", err)
		return
	}

	fmt.Println(guild)

	// h.session.ChannelMessageSend(guild.SystemChannelID, "Hi, Me triggered every minutes")
}

func (h *cronJobHandler) testFiveMinuteJob() {
	fmt.Println("Every 5 minutes")
}
