package groupmestatsbot

import (
	"strings"

	"github.com/MagnusFrater/groupme"
)

const messageDivider = "==============================" // 30 '='

// Stats contains a GroupMe group's statistics.
type Stats struct {
	Messages            []*groupme.Message  // GroupMe Messages to analyze
	Members             map[string]*Member  // UserID -> *Member
	WordFrequency       map[string]*Word    // text -> *Word
	CharacterFrequency  map[rune]*Character // rune -> *Character
	Reposts             map[string]*Repost  // text -> *Repost
	TotalMessagesLength int                 // the length of all messages combined together

	BlacklistedUserIDs map[string]struct{} // UserIDs to ignore while analyzing messages; UserID -> nil
}

// NewStats creates a new Stats.
func NewStats(messages []*groupme.Message) Stats {
	return Stats{
		Messages:           messages,
		Members:            make(map[string]*Member),
		WordFrequency:      make(map[string]*Word),
		CharacterFrequency: make(map[rune]*Character),
		Reposts:            make(map[string]*Repost),

		BlacklistedUserIDs: make(map[string]struct{}),
	}
}

// Analyze analyzes a GroupMe group's messages.
func (s *Stats) Analyze() {
	for _, message := range s.Messages {
		// don't analyze blacklisted users
		if s.Blacklisted(message.UserID) {
			s.addMember(message.UserID, message.Name) // just in case
			continue
		}

		// parse reposts
		s.incRepost(message)

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

		// parse visionary
		if len(message.Attachments) > 0 {
			for _, attachment := range message.Attachments {
				switch attachment.Type {
				case groupme.ImageAttachment:
					s.incVisionary(message.UserID, message.Name)
				}
			}
		}
	}
}

// Blacklist blacklists a UserID such that it is ignored while analyzing messages.
func (s *Stats) Blacklist(userID string) {
	if _, ok := s.BlacklistedUserIDs[userID]; !ok {
		s.BlacklistedUserIDs[userID] = struct{}{}
	}
}

// Blacklisted returns whether the given UserID is blacklisted from being analyzed.
func (s *Stats) Blacklisted(userID string) bool {
	_, ok := s.BlacklistedUserIDs[userID]
	return ok
}
