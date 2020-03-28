package groupmestatsbot

import (
	"fmt"
	"sort"
)

// Word is a single text token. These make up a GroupMe message's text.
type Word struct {
	Text      string
	Frequency int
}

func (s *Stats) addWord(text string) {
	if _, ok := s.WordFrequency[text]; !ok {
		s.WordFrequency[text] = &Word{Text: text}
	}
}

func (s *Stats) incWord(text string) {
	s.addWord(text)

	s.WordFrequency[text].Frequency++
}

// TopWords returns a sorted list of the most frequently used words.
func (s *Stats) TopWords(limit int) []*Word {
	sorted := []*Word{}

	for _, w := range s.WordFrequency {
		sorted = append(sorted, w)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Frequency > sorted[j].Frequency })

	top := []*Word{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

// SprintTopWords formats a Top Words Bot post and returns the resulting string.
func (s *Stats) SprintTopWords(limit int) string {
	str := fmt.Sprintf("Top Words\n%s\n", messageDivider)

	topWords := s.TopWords(limit)
	for i, w := range topWords {
		str += fmt.Sprintf("%d) %s: %d", i+1, w.Text, w.Frequency)

		// don't put newline after last ranking
		if i < len(topWords)-1 {
			str += "\n"
		}
	}

	return str
}
