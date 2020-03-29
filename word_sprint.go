package groupmestatsbot

import (
	"fmt"
)

// SprintTopWords formats a Top Words Bot post and returns the resulting string.
func (s *Stats) SprintTopWords(limit int) string {
	str := fmt.Sprintf("Top Words\n%s\n", messageDivider)

	topWords := s.TopWords(limit)
	if len(topWords) == 0 {
		str += "\nThere are no words."
		return str
	}

	for i, w := range topWords {
		str += fmt.Sprintf("%d) %s: %d", i+1, w.Text, w.Frequency)

		// don't put newline after last ranking
		if i < len(topWords)-1 {
			str += "\n"
		}
	}

	return str
}
