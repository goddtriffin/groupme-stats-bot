package groupmestatsbot

import (
	"errors"
	"log"

	"github.com/MagnusFrater/groupme"
)

// Commands
const (
	// all
	CommandAll = "all"

	// members
	CommandTopOfThePops        = "topOfThePops"
	CommandTopOfTheSimps       = "topOfTheSimps"
	CommandTopOfTheNarcissists = "topOfTheNarcissists"
	CommandTopPoster           = "topPoster"
	CommandMostCharismatic     = "mostCharismatic"
	CommandTopLurker           = "topLurker"
	CommandTopRambler          = "topRambler"
	CommandMostVisionary       = "mostVisionary"
	CommandTopWordsmith        = "topWordsmith"
	CommandBiggestFoot         = "biggestFoot"
	CommandSorestBum           = "sorestBum"
	CommandTopMother           = "topMother"
	CommandMostReincarnated    = "mostReincarnated"

	// messages
	CommandTextFrequencyAnalysis = "textFrequencyAnalysis"
	CommandTopMessages           = "topMessages"
	CommandTopReposts            = "topReposts"
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

// Command runs a GroupMe Stats Bot command
// The returned bool represents whether or not a command was ran.
func (b *StatsBot) Command(command string, logOnly bool) (bool, error) {
	var err error
	var output string

	switch command {
	case CommandAll:
		err = b.AllCommands(logOnly)
	case CommandTopOfThePops:
		output = b.Stats.SprintTopOfThePops(b.Limit)
	case CommandTopOfTheSimps:
		output = b.Stats.SprintTopOfTheSimps(b.Limit)
	case CommandTopOfTheNarcissists:
		output = b.Stats.SprintTopOfTheNarcissists(b.Limit)
	case CommandTopPoster:
		output = b.Stats.SprintTopPoster(b.Limit)
	case CommandMostCharismatic:
		output = b.Stats.SprintMostCharismatic(b.Limit)
	case CommandTopLurker:
		output = b.Stats.SprintTopLurker(b.Limit)
	case CommandTopRambler:
		output = b.Stats.SprintTopRambler(b.Limit)
	case CommandMostVisionary:
		output = b.Stats.SprintMostVisionary(b.Limit)
	case CommandTopWordsmith:
		output = b.Stats.SprintTopWordsmith(b.Limit)
	case CommandBiggestFoot:
		output = b.Stats.SprintBiggestFoot(b.Limit)
	case CommandSorestBum:
		output = b.Stats.SprintSorestBum(b.Limit)
	case CommandTopMother:
		output = b.Stats.SprintTopMother(b.Limit)
	case CommandMostReincarnated:
		output = b.Stats.SprintMostReincarnated(b.Limit)
	case CommandTextFrequencyAnalysis:
		output = b.Stats.SprintTextFrequencyAnalysis(b.Limit)
	case CommandTopMessages:
		output = b.Stats.SprintTopMessages(b.Limit)
	case CommandTopReposts:
		output = b.Stats.SprintTopReposts(b.Limit)
	}

	// no command was ran
	if command != CommandAll && output == "" {
		return false, nil
	}

	if logOnly {
		log.Printf("\n%s\n", output)
	} else {
		err = b.Bot.Post(output, nil)
		if err != nil {
			return true, err
		}
	}

	return true, nil
}

// AllCommands runs every single GroupMe Stats Bot command.
func (b *StatsBot) AllCommands(logOnly bool) error {
	commands := GetAllCommands()

	var err error
	for _, command := range commands {
		_, err = b.Command(command, logOnly)
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
		CommandTopOfThePops, CommandTopOfTheSimps, CommandTopOfTheNarcissists,
		CommandTopPoster, CommandMostCharismatic, CommandTopLurker,
		CommandTopRambler, CommandMostVisionary, CommandTopWordsmith,
		CommandBiggestFoot, CommandSorestBum, CommandTopMother,
		CommandMostReincarnated,
		// messages
		CommandTextFrequencyAnalysis, CommandTopMessages, CommandTopReposts,
	}
}
