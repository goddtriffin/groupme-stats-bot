package groupmestatsbot

import (
	"errors"

	"github.com/MagnusFrater/groupme"
)

// Commands
const (
	// all
	All = "all"

	// members
	TopOfThePops        = "topOfThePops"
	TopOfTheSimps       = "topOfTheSimps"
	TopOfTheNarcissists = "topOfTheNarcissists"
	TopPoster           = "topPoster"
	MostCharismatic     = "mostCharismatic"
	TopLurker           = "topLurker"
	TopRambler          = "topRambler"
	MostVisionary       = "mostVisionary"
	TopWordsmith        = "topWordsmith"
	BiggestFoot         = "biggestFoot"
	SorestBum           = "sorestBum"

	// messages
	TextFrequencyAnalysis = "textFrequencyAnalysis"
	TopMessages           = "topMessages"
	TopReposts            = "topReposts"
)

// StatsBot is the GroupMe Stats Bot.
type StatsBot struct {
	Client groupme.Client
	Bot    groupme.Bot

	Stats *Stats
	Limit int
}

// New returns a new StatsBot.
func New(accessToken, botID, groupID string, limit int, blacklist []string) (StatsBot, error) {
	statsBot := StatsBot{}

	if accessToken == "" {
		return statsBot, errors.New("invalid accessToken")
	}
	if botID == "" {
		return statsBot, errors.New("invalid botID")
	}
	if groupID == "" {
		return statsBot, errors.New("invalid groupID")
	}

	// limit should always be positive
	if limit < 1 {
		limit = 5
	}
	statsBot.Limit = limit

	statsBot.Client = groupme.NewClient(groupme.V3BaseURL, accessToken)
	statsBot.Bot = groupme.NewBot(groupme.V3BaseURL, botID, groupID, "", "")

	messages, err := statsBot.Client.AllMessages(statsBot.Bot.GroupID)
	if err != nil {
		return statsBot, err
	}
	statsBot.Stats = NewStats(messages)

	if len(blacklist) != 0 {
		for _, ID := range blacklist {
			statsBot.Stats.Blacklist(ID)
		}
	}

	statsBot.Stats.Analyze()

	return statsBot, nil
}

// Command runs a GroupMe Stats Bot command.
func (b *StatsBot) Command(command string) (bool, error) {
	var err error
	triedCommand := true

	switch command {
	case All:
		err = b.AllCommands()
	case TopOfThePops:
		err = b.Bot.Post(b.Stats.SprintTopOfThePops(b.Limit), nil)
	case TopOfTheSimps:
		err = b.Bot.Post(b.Stats.SprintTopOfTheSimps(b.Limit), nil)
	case TopOfTheNarcissists:
		err = b.Bot.Post(b.Stats.SprintTopOfTheNarcissists(b.Limit), nil)
	case TopPoster:
		err = b.Bot.Post(b.Stats.SprintTopPoster(b.Limit), nil)
	case MostCharismatic:
		err = b.Bot.Post(b.Stats.SprintMostCharismatic(b.Limit), nil)
	case TopLurker:
		err = b.Bot.Post(b.Stats.SprintTopLurker(b.Limit), nil)
	case TopRambler:
		err = b.Bot.Post(b.Stats.SprintTopRambler(b.Limit), nil)
	case MostVisionary:
		err = b.Bot.Post(b.Stats.SprintMostVisionary(b.Limit), nil)
	case TopWordsmith:
		err = b.Bot.Post(b.Stats.SprintTopWordsmith(b.Limit), nil)
	case BiggestFoot:
		err = b.Bot.Post(b.Stats.SprintBiggestFoot(b.Limit), nil)
	case SorestBum:
		err = b.Bot.Post(b.Stats.SprintSorestBum(b.Limit), nil)
	case TextFrequencyAnalysis:
		err = b.Bot.Post(b.Stats.SprintTextFrequencyAnalysis(b.Limit), nil)
	case TopMessages:
		err = b.Bot.Post(b.Stats.SprintTopMessages(b.Limit), nil)
	case TopReposts:
		err = b.Bot.Post(b.Stats.SprintTopReposts(b.Limit), nil)
	default:
		triedCommand = false
	}

	if err != nil {
		return triedCommand, err
	}

	return triedCommand, nil
}

// AllCommands runs every single GroupMe Stats Bot command.
func (b *StatsBot) AllCommands() error {
	commands := GetAllCommands()

	var err error
	for _, command := range commands {
		_, err = b.Command(command)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetAllCommands returns a list of all available commands.
func GetAllCommands() []string {
	return []string{
		// members
		TopOfThePops, TopOfTheSimps, TopOfTheNarcissists,
		TopPoster, MostCharismatic, TopLurker,
		TopRambler, MostVisionary, TopWordsmith,
		BiggestFoot, SorestBum,
		// messages
		TextFrequencyAnalysis, TopMessages, TopReposts,
	}
}
