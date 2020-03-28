package groupmestatsbot

import "fmt"

// SprintTextFrequencyAnalysis formats a Text Frequency Analysis Bot post and returns the resulting string.
func (s *Stats) SprintTextFrequencyAnalysis(limit int) string {
	str := "KOWALSKI, ANALYSIS!\n\n"

	str += fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s",
		s.SprintTotalMessages(),
		s.SprintAverageMessageLength(),
		s.SprintTopWords(limit),
		s.SprintTopCharacters(limit))

	return str
}
