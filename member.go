package groupmestatsbot

import (
	"fmt"
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

// SprintTopOfThePops formats a Top of the Pops Bot post and returns the resulting string.
func (s *Stats) SprintTopOfThePops(limit int) string {
	str := "Top of the Pops\n(who has the most upvotes)\n==========\n"

	topPopulars := s.TopOfThePops(limit)
	for i, member := range topPopulars {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.PopularityScore)

		// don't put newline after last ranking
		if i < len(topPopulars)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintTopOfTheSimps formats a Top of the Simps Bot post and returns the resulting string.
func (s *Stats) SprintTopOfTheSimps(limit int) string {
	str := "Top of the Simps\n(who upvoted other people the most)\n==========\n"

	topSimps := s.TopOfTheSimps(limit)
	for i, member := range topSimps {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.SimpScore)

		// don't put newline after last ranking
		if i < len(topSimps)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintTopOfTheNarcissists formats a Top of the Narcissists Bot post and returns the resulting string.
func (s *Stats) SprintTopOfTheNarcissists(limit int) string {
	str := "Top of the Narcissists\n(who upvoted themselves the most)\n==========\n"

	topNarcissists := s.TopOfTheNarcissists(limit)
	for i, member := range topNarcissists {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.NarcissistScore)

		// don't put newline after last ranking
		if i < len(topNarcissists)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintTopPoster formats a Top Poster Bot post and returns the resulting string.
func (s *Stats) SprintTopPoster(limit int) string {
	str := "Top Poster\n(who posted the most)\n==========\n"

	topPosters := s.TopPosters(limit)
	for i, member := range topPosters {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.NumMessages)

		// don't put newline after last ranking
		if i < len(topPosters)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintMostCharismatic formats a Most Charismatic Bot post and returns the resulting string.
func (s *Stats) SprintMostCharismatic(limit int) string {
	str := "Most Charismatic\n(# of likes / # of messages)\n==========\n"

	mostCharismatic := s.MostCharismatic(limit)
	for i, member := range mostCharismatic {
		str += fmt.Sprintf("%d) %s: %.3f", i+1, member.Name, member.Charisma())

		// don't put newline after last ranking
		if i < len(mostCharismatic)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintTopLurker formats a Top Lurker Bot post and returns the resulting string.
func (s *Stats) SprintTopLurker(limit int) string {
	str := "Top Lurker\n(# of likes given / # of messages)\n==========\n"

	topLurker := s.TopLurker(limit)
	for i, member := range topLurker {
		str += fmt.Sprintf("%d) %s: %.3f", i+1, member.Name, member.Lurky())

		// don't put newline after last ranking
		if i < len(topLurker)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintTopRambler formats a Top Rambler Bot post and returns the resulting string.
func (s *Stats) SprintTopRambler(limit int) string {
	str := "Top Rambler\n(most messages with zero likes)\n==========\n"

	topRambler := s.TopRambler(limit)
	for i, member := range topRambler {
		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.UnpopularityScore)

		// don't put newline after last ranking
		if i < len(topRambler)-1 {
			str += "\n"
		}
	}

	return str
}
