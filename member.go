package groupmestatsbot

import (
	"fmt"
	"sort"
)

// Member is a container for a GroupMe member's staistics.
type Member struct {
	ID              string
	Name            string
	PopularityScore int // how often did others upvote them
	SimpScore       int // how many times did they upvote someone else
	NarcissistScore int // how many times did they upvote themselves
	NumMessages     int // how many messages did they send
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

func (s *Stats) incPopularity(userID, name string, inc int) {
	s.addMember(userID, name)

	s.Members[userID].PopularityScore += inc
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
