package groupmestatsbot

import (
	"sort"
)

// Member is a container for a GroupMe member's staistics.
type Member struct {
	ID                string
	Name              string
	PopularityScore   int // how often did others upvote them
	UnpopularityScore int // how often did their messages get zero favorites
	SimpScore         int // how many times did they upvote someone else
	NarcissistScore   int // how many times did they upvote themselves
	NumMessages       int // how many messages did they send
}

// Charisma returns the quality of their overall messages.
func (m *Member) Charisma() float64 {
	if m.PopularityScore < 1 || m.NumMessages < 1 {
		return -1
	}

	return float64(m.PopularityScore) / float64(m.NumMessages)
}

// Lurky returns a ratio between their interactions with others' messages and how often they post messages themselves.
func (m *Member) Lurky() float64 {
	if m.SimpScore < 1 || m.NumMessages < 1 {
		return -1
	}

	return float64(m.SimpScore) / float64(m.NumMessages)
}

func (s *Stats) addMember(userID, name string) {
	if m, ok := s.Members[userID]; !ok {
		s.Members[userID] = &Member{
			ID:   userID,
			Name: name,
		}
	} else {
		if m.Name == "" {
			m.Name = name
		}
	}
}

func (s *Stats) incNumMessages(userID, name string) {
	s.addMember(userID, name)

	s.Members[userID].NumMessages++
}

func (s *Stats) incPopularity(userID, name string) {
	s.addMember(userID, name)

	s.Members[userID].PopularityScore++
}

func (s *Stats) incUnpopularity(userID, name string) {
	s.addMember(userID, name)

	s.Members[userID].UnpopularityScore++
}

func (s *Stats) incSimp(userID, name string) {
	s.addMember(userID, name)

	s.Members[userID].SimpScore++
}

func (s *Stats) incNarcissist(userID, name string) {
	s.addMember(userID, name)

	s.Members[userID].NarcissistScore++
}

// TopOfThePops returns a sorted list of the most popular members.
// Popularity is defined as who has the most upvotes.
func (s *Stats) TopOfThePops(limit int) []*Member {
	sorted := []*Member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].PopularityScore > sorted[j].PopularityScore })

	top := []*Member{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

// TopOfTheSimps returns a sorted list of the biggest simp members.
// A simp is defined as who upvotes other members the most.
func (s *Stats) TopOfTheSimps(limit int) []*Member {
	sorted := []*Member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].SimpScore > sorted[j].SimpScore })

	top := []*Member{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

// TopOfTheNarcissists returns a sorted list of the biggest narcissistic members.
// A narcissist is defined as who upvotes themselves the most.
func (s *Stats) TopOfTheNarcissists(limit int) []*Member {
	sorted := []*Member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].NarcissistScore > sorted[j].NarcissistScore })

	top := []*Member{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

// TopPosters returns a sorted list of who posted the most messages.
func (s *Stats) TopPosters(limit int) []*Member {
	sorted := []*Member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].NumMessages > sorted[j].NumMessages })

	top := []*Member{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

// MostCharismatic returns a sorted list of who posts the highest quality messages.
// Charisma is defined as # of likes / # of messages.
func (s *Stats) MostCharismatic(limit int) []*Member {
	sorted := []*Member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Charisma() > sorted[j].Charisma() })

	top := []*Member{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

// TopLurker returns a sorted list of who lurks the most.
// A lurker is defined as # of likes given out / # of messages posted.
func (s *Stats) TopLurker(limit int) []*Member {
	sorted := []*Member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Lurky() > sorted[j].Lurky() })

	top := []*Member{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

// TopRambler returns a sorted list of who has the most messages with zero favorites.
func (s *Stats) TopRambler(limit int) []*Member {
	sorted := []*Member{}

	for _, member := range s.Members {
		sorted = append(sorted, member)
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].UnpopularityScore > sorted[j].UnpopularityScore })

	top := []*Member{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}
