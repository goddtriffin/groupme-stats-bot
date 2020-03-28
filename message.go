package groupmestatsbot

import (
	"fmt"
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

// SprintTotalMessages formats an Total Messages Bot post and returns the resulting string.
func (s *Stats) SprintTotalMessages() string {
	return fmt.Sprintf("Total Messages: %d messages", s.TotalMessages())
}

// SprintAverageMessageLength formats an Average Message Length Bot post and returns the resulting string.
func (s *Stats) SprintAverageMessageLength() string {
	return fmt.Sprintf("Average Message Length: %d words", s.AverageMessageLength())
}

// SprintTopMessages formats a Top Messages Bot post and returns the resulting string.
func (s *Stats) SprintTopMessages(limit int) string {
	str := fmt.Sprintf("Top Messages\n%s\n", messageDivider)

	topMessages := s.TopMessages(limit)
	for i, message := range topMessages {
		str += fmt.Sprintf("%d) (%d) %s:\n", i+1, len(message.FavoritedBy), message.Name)

		if len(message.Attachments) > 0 {
			for i, attachment := range message.Attachments {
				switch attachment.Type {
				case groupme.ImageAttachment:
					str += fmt.Sprintf("image: %s", attachment.URL)
				}

				// online put newline if there are more attachments, or if there is message text
				if i < len(message.Attachments)-1 || message.Text != "" {
					str += "\n"
				}
			}
		}

		if message.Text != "" {
			str += fmt.Sprintf("\"%s\"", message.Text)
		}

		// don't put newline after last ranking
		if i < len(topMessages)-1 {
			str += "\n\n"
		}
	}

	return str
}
