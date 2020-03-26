package groupmestatsbot

import (
	"strings"

	"github.com/MagnusFrater/groupme"
)

// Stats contains a GroupMe group's statistics.
type Stats struct {
	Messages           []groupme.Message
	Members            map[string]*Member
	WordFrequency      map[string]*Word
	CharacterFrequency map[rune]*Character
}

// NewStats creates a new Stats.
func NewStats(messages []groupme.Message) Stats {
	return Stats{
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
		s.incPopularity(message.UserID, message.Name, len(message.FavoritedBy))

		// parse narcissists and simps
		for _, userID := range message.FavoritedBy {
			if userID == message.UserID {
				s.incNarcissist(message.UserID, message.Name)
			} else {
				s.incSimp(userID, "")
			}
		}

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
