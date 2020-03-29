package groupmestatsbot

import (
	"sort"
)

// Character is a single text character.
type Character struct {
	R         rune
	Frequency int
}

func (s *Stats) addCharacter(r rune) {
	if _, ok := s.CharacterFrequency[r]; !ok {
		s.CharacterFrequency[r] = &Character{R: r}
	}
}

func (s *Stats) incCharacter(r rune) {
	s.addCharacter(r)

	s.CharacterFrequency[r].Frequency++
}

// TopCharacters returns a sorted list of the most frequently used characters.
func (s *Stats) TopCharacters(limit int) []*Character {
	sorted := []*Character{}

	for _, c := range s.CharacterFrequency {
		sorted = append(sorted, c)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Frequency > sorted[j].Frequency })

	top := []*Character{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}
