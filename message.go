package groupmestatsbot

import (
	"sort"

	"github.com/MagnusFrater/groupme"
)

// TotalMessages returns the total number of messages.
func (s *Stats) TotalMessages() int {
	return len(s.Messages)
}

// AverageMessageLength returns the average message length.
func (s *Stats) AverageMessageLength() int {
	return s.TotalMessagesLength / len(s.Messages)
}

// TopMessages returns a sorted list of the most favorited messages.
func (s *Stats) TopMessages(limit int) []*groupme.Message {
	sorted := make([]*groupme.Message, len(s.Messages))
	copy(sorted, s.Messages)

	sort.Slice(sorted, func(i, j int) bool { return len(sorted[i].FavoritedBy) > len(sorted[j].FavoritedBy) })

	top := []*groupme.Message{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}
