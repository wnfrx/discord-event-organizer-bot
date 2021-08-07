package command

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/wnfrx/discord-event-organizer-bot/models"
)

func (h *botCommandHandler) commandHandlerVoting(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Options[0].Name {
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Oops, something gone wrong.\n" +
					"Hol' up, you aren't supposed to see this message.",
			},
		})

	case CommandVoteCreate:
		h.subcommandHandlerVotingCreate(s, i)
	}
}

func (h *botCommandHandler) subcommandHandlerVotingCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Printf("param: %+v", i.ApplicationCommandData().Options[0].Options[0].StringValue())
	log.Printf("param: %+v", i.ApplicationCommandData().Options[0].Options[1].IntValue())

	description := i.ApplicationCommandData().Options[0].Options[0].StringValue()
	duration := i.ApplicationCommandData().Options[0].Options[1].IntValue()

	log.Println("create initial vote")

	var (
		yEmoji      = "👍"
		nEmoji      = "👎"
		cancelEmoji = "🚫"
		isCanceled  bool
	)

	vote := models.Voting{
		Name:      description,
		GuildID:   i.GuildID,
		ChannelID: i.ChannelID,
		Options: []models.VotingOption{
			{
				Name: yEmoji,
			},
			{
				Name: nEmoji,
			},
		},
	}

	h.votingGuildMap[i.GuildID] = vote

	defer func() {
		delete(h.votingGuildMap, i.GuildID)
	}()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Voting dimulai!",
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: description,
					Description: `
						Yep: 0
						Nop: 0
					`,
				},
			},
		},
	})

	interactionMsg, err := s.InteractionResponse(s.State.User.ID, i.Interaction)
	if err != nil {
		return
	}

	// Add voting emojis
	if err = s.MessageReactionAdd(i.ChannelID, interactionMsg.ID, yEmoji); err != nil {
		log.Println(err)
		return
	}
	if err = s.MessageReactionAdd(i.ChannelID, interactionMsg.ID, nEmoji); err != nil {
		log.Println(err)
		return
	}
	if err = s.MessageReactionAdd(i.ChannelID, interactionMsg.ID, cancelEmoji); err != nil {
		log.Println(err)
		return
	}

	// Voting process
	ticker := time.NewTicker(time.Duration(duration) * time.Second)
	defer ticker.Stop()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ticker.C:
				// voting finished
				return

			default:
				cancelUsers, err := s.MessageReactions(i.ChannelID, interactionMsg.ID, cancelEmoji, 100, "", "")
				if err != nil {
					continue
				}

				yUsers, err := s.MessageReactions(i.ChannelID, interactionMsg.ID, yEmoji, 100, "", "")
				if err != nil {
					continue
				}

				nUsers, err := s.MessageReactions(i.ChannelID, interactionMsg.ID, nEmoji, 100, "", "")
				if err != nil {
					continue
				}

				yOption := models.VotingOption{
					Name: yEmoji,
				}

				nOption := models.VotingOption{
					Name: nEmoji,
				}

				for _, u := range yUsers {
					yOption.Users = append(yOption.Users, u.ID)
				}

				for _, u := range nUsers {
					nOption.Users = append(nOption.Users, u.ID)
				}

				vote.Options = []models.VotingOption{yOption, nOption}
				h.votingGuildMap[i.GuildID] = vote

				if len(cancelUsers) > 1 {
					log.Println("voting canceled")
					isCanceled = true

					err = s.FollowupMessageEdit(s.State.User.ID, i.Interaction, interactionMsg.ID, &discordgo.WebhookEdit{
						Content: "Voting batal gan",
						Embeds: []*discordgo.MessageEmbed{
							{
								Title: description,
								Description: fmt.Sprintf(
									`
									Yep: %d
									Nop: %d
								`,
									yOption.CountVote(),
									nOption.CountVote(),
								),
							},
						},
					})

					return
				}

				err = s.FollowupMessageEdit(s.State.User.ID, i.Interaction, interactionMsg.ID, &discordgo.WebhookEdit{
					Embeds: []*discordgo.MessageEmbed{
						{
							Title: description,
							Description: fmt.Sprintf(
								`
									Yep: %d
									Nop: %d
								`,
								yOption.CountVote(),
								nOption.CountVote(),
							),
						},
					},
				})
			}
		}
	}()

	wg.Wait()

	if isCanceled {
		return
	}

	// decision
	var (
		totalVote, yCount, nCount int64
		decisionContent           string
	)

	totalVote = vote.CountTotalVotes()
	yCount = vote.Options[0].CountVote()
	nCount = vote.Options[1].CountVote()

	switch {
	case yCount > nCount:
		decisionContent = fmt.Sprintf("Voting kelar. Yep! (%d/%d)", yCount, totalVote)
	case nCount > yCount:
		decisionContent = fmt.Sprintf("Voting kelar. Nope! (%d/%d)", nCount, totalVote)
	default:
		decisionContent = "Voting kelar. Seri!"
	}

	err = s.FollowupMessageEdit(s.State.User.ID, i.Interaction, interactionMsg.ID, &discordgo.WebhookEdit{
		Content: decisionContent,
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: description,
				Description: fmt.Sprintf(
					`
						Yep: %d
						Nop: %d
					`,
					yCount,
					nCount,
				),
			},
		},
	})

}
