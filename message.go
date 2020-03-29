package groupmestatsbot

import (
	"math"
	"sort"

	"github.com/MagnusFrater/groupme"
)

// TotalMessages returns the total number of messages.
func (s *Stats) TotalMessages() int {
	return len(s.Messages)
}

// AverageMessageLength returns the average message length.
func (s *Stats) AverageMessageLength() int {
	if s.TotalMessagesLength == 0 || len(s.Messages) == 0 {
		return -1
	}

	return s.TotalMessagesLength / len(s.Messages)
}

// TopMessages returns a sorted list of the most favorited messages.
func (s *Stats) TopMessages(limit int) []*groupme.Message {
	if limit == -1 {
		limit = math.MaxInt64
	}

	sorted := make([]*groupme.Message, len(s.Messages))
	copy(sorted, s.Messages)

	sort.Slice(sorted, func(i, j int) bool { return len(sorted[i].FavoritedBy) > len(sorted[j].FavoritedBy) })

	top := []*groupme.Message{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}
