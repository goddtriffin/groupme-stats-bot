package groupmestatsbot

import (
	"fmt"
)

// SprintTopOfThePops formats a Top of the Pops Bot post and returns the resulting string.
func (s *Stats) SprintTopOfThePops(limit int) string {
	str := fmt.Sprintf("Top of the Pops\n(who has the most favorites)\n%s\n", messageDivider)

	topPopulars := s.TopOfThePops(limit)
	if len(topPopulars) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range topPopulars {
		if member.PopularityScore == 0 {
			str += "\nEveryone else is unpopular."
			break
		}

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
	str := fmt.Sprintf("Top of the Simps\n(who favorited other people's messages the most)\n%s\n", messageDivider)

	topSimps := s.TopOfTheSimps(limit)
	if len(topSimps) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range topSimps {
		if member.SimpScore == 0 {
			str += "\nEveryone else is not a simp."
			break
		}

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
	str := fmt.Sprintf("Top of the Narcissists\n(who favorited their own messages the most)\n%s\n", messageDivider)

	topNarcissists := s.TopOfTheNarcissists(limit)
	if len(topNarcissists) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range topNarcissists {
		if member.NarcissistScore == 0 {
			str += "\nEveryone else is not a narcissist."
			break
		}

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
	str := fmt.Sprintf("Top Poster\n(who posted the most messages)\n%s\n", messageDivider)

	topPosters := s.TopPosters(limit)
	if len(topPosters) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range topPosters {
		if member.NumMessages == 0 {
			str += "\nEveryone else posted 0 messages."
			break
		}

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
	str := fmt.Sprintf("Most Charismatic\n(# of favorites received / # of messages they posted)\n%s\n", messageDivider)

	mostCharismatic := s.MostCharismatic(limit)
	if len(mostCharismatic) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range mostCharismatic {
		if member.Charisma() == -1 {
			str += "\nEveryone else has either received 0 favorites on their messages, or has posted 0 messages themselves."
			break
		}

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
	str := fmt.Sprintf("Top Lurker\n(# of favorites given out / # of messages they posted)\n%s\n", messageDivider)

	topLurker := s.TopLurker(limit)
	if len(topLurker) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range topLurker {
		if member.Lurkiness() == -1 {
			str += "\nEveryone else either hasn't given out any favorites, or has posted 0 messages themselves."
			break
		}

		str += fmt.Sprintf("%d) %s: %.3f", i+1, member.Name, member.Lurkiness())

		// don't put newline after last ranking
		if i < len(topLurker)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintTopRambler formats a Top Rambler Bot post and returns the resulting string.
func (s *Stats) SprintTopRambler(limit int) string {
	str := fmt.Sprintf("Top Rambler\n(who posted the most messages with zero favorites)\n%s\n", messageDivider)

	topRambler := s.TopRambler(limit)
	if len(topRambler) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range topRambler {
		if member.UnpopularityScore == 0 {
			str += "\nEveryone else has favorites on all of their messages."
			break
		}

		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.UnpopularityScore)

		// don't put newline after last ranking
		if i < len(topRambler)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintMostVisionary formats a Most Visionary Bot post and returns the resulting string.
func (s *Stats) SprintMostVisionary(limit int) string {
	str := fmt.Sprintf("Most Visionary\n(who posted the most images)\n%s\n", messageDivider)

	mostVisionary := s.MostVisionary(limit)
	if len(mostVisionary) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range mostVisionary {
		if member.VisionaryScore == 0 {
			str += "\nEveryone else posted 0 images."
			break
		}

		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.VisionaryScore)

		// don't put newline after last ranking
		if i < len(mostVisionary)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintTopWordsmith formats a Top Wordsmith Bot post and returns the resulting string.
func (s *Stats) SprintTopWordsmith(limit int) string {
	str := fmt.Sprintf("Top Wordsmith\n(who posted the most text-only messages)\n%s\n", messageDivider)

	topWordsmith := s.TopWordsmith(limit)
	if len(topWordsmith) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range topWordsmith {
		if member.WordsmithScore == 0 {
			str += "\nEveryone else posted messages with no text (aka messages with attachments)."
			break
		}

		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, member.WordsmithScore)

		// don't put newline after last ranking
		if i < len(topWordsmith)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintBiggestFoot formats a Biggest Foot Bot post and returns the resulting string.
func (s *Stats) SprintBiggestFoot(limit int) string {
	str := fmt.Sprintf("Biggest Foot\n(who kicked the most members from the group)\n%s\n", messageDivider)

	biggestFoot := s.BiggestFoot(limit)
	if len(biggestFoot) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range biggestFoot {
		if len(member.Kicked) == 0 {
			str += "\nEveryone else kicked 0 people."
			break
		}

		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, len(member.Kicked))

		// don't put newline after last ranking
		if i < len(biggestFoot)-1 {
			str += "\n"
		}
	}

	return str
}

// SprintSorestBum formats a Sorest Bum Bot post and returns the resulting string.
func (s *Stats) SprintSorestBum(limit int) string {
	str := fmt.Sprintf("Sorest Bum\n(who got kicked from the group the most)\n%s\n", messageDivider)

	sorestBum := s.SorestBum(limit)
	if len(sorestBum) == 0 {
		str += "\nThere are no members."
		return str
	}

	for i, member := range sorestBum {
		if len(member.KickedBy) == 0 {
			str += "\nEveryone else was kicked 0 times."
			break
		}

		str += fmt.Sprintf("%d) %s: %d", i+1, member.Name, len(member.KickedBy))

		// don't put newline after last ranking
		if i < len(sorestBum)-1 {
			str += "\n"
		}
	}

	return str
}
