package groupmestatsbot

import (
	"fmt"
)

// SprintTopCharacters formats a Top Characters Bot post and returns the resulting string.
func (s *Stats) SprintTopCharacters(limit int) string {
	str := fmt.Sprintf("Top Characters\n%s\n", messageDivider)

	topCharacters := s.TopCharacters(limit)
	for i, c := range topCharacters {
		str += fmt.Sprintf("%d) %s: %d", i+1, string(c.R), c.Frequency)

		// don't put newline after last ranking
		if i < len(topCharacters)-1 {
			str += "\n"
		}
	}

	return str
}
