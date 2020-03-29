package groupmestatsbot

import (
	"fmt"
)

// SprintTopOfThePops formats a Top of the Pops Bot post and returns the resulting string.
func (s *Stats) SprintTopOfThePops(limit int) string {
	str := fmt.Sprintf("Top of the Pops\n(who has the most upvotes)\n%s\n", messageDivider)

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
	str := fmt.Sprintf("Top of the Simps\n(who upvoted other people the most)\n%s\n", messageDivider)

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
	str := fmt.Sprintf("Top of the Narcissists\n(who upvoted themselves the most)\n%s\n", messageDivider)

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
	str := fmt.Sprintf("Top Poster\n(who posted the most)\n%s\n", messageDivider)

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
	str := fmt.Sprintf("Most Charismatic\n(# of likes / # of messages)\n%s\n", messageDivider)

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
	str := fmt.Sprintf("Top Lurker\n(# of likes given / # of messages)\n%s\n", messageDivider)

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
	str := fmt.Sprintf("Top Rambler\n(most messages with zero likes)\n%s\n", messageDivider)

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
