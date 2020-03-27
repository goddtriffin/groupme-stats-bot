package groupmestatsbot

import (
	"strings"

	"github.com/MagnusFrater/groupme"
)

// Stats contains a GroupMe group's statistics.
type Stats struct {
	Messages            []groupme.Message
	Members             map[string]*Member
	WordFrequency       map[string]*Word
	CharacterFrequency  map[rune]*Character
	TotalMessagesLength int // the length of all messages combined together
}

// NewStats creates a new Stats.
func NewStats(messages []groupme.Message) Stats {
	return Stats{
		Messages:           messages,
		Members:            make(map[string]*Member),
		CharacterFrequency: make(map[rune]*Character),
		WordFrequency:      make(map[string]*Word),
	}
}

// Analyze analyzes a GroupMe group's messages.
func (s *Stats) Analyze() {
	for _, message := range s.Messages {
		// parse numMessage and popularity
		s.incNumMessages(message.UserID, message.Name)

		if len(message.FavoritedBy) > 0 {
			// parse narcissists and simps
			for _, userID := range message.FavoritedBy {
				if userID == message.UserID {
					s.incNarcissist(message.UserID, message.Name)
				} else {
					s.incPopularity(message.UserID, message.Name)
					s.incSimp(userID, "")
				}
			}
		} else {
			// unpopularity - their message received zero favorites
			s.incUnpopularity(message.UserID, message.Name)
		}

		// parse message length
		s.TotalMessagesLength += len(message.Text)

		// parse word frequency
		for _, text := range strings.Fields(message.Text) {
			s.incWord(text)

			// parse character frequency
			runes := []rune(text)
			for _, r := range runes {
				s.incCharacter(r)
			}
		}
	}
}
