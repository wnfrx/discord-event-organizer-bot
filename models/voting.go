package models

type Voting struct {
	Name      string
	GuildID   string
	ChannelID string
	Options   []VotingOption
}

type VotingOption struct {
	Name  string
	Users []string
}

func (v Voting) CountTotalVotes() int64 {
	var total int64

	for _, opt := range v.Options {
		total += int64(len(opt.Users))
	}

	return total
}

func (o *VotingOption) Add(userID string) {
	o.Users = append(o.Users, userID)
	return
}

func (o VotingOption) CountVote() int64 {
	return int64(len(o.Users))
}
