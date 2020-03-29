package groupmestatsbot

import (
	"math"
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
	if limit == -1 {
		limit = math.MaxInt64
	}

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
