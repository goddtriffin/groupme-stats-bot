package groupmestatsbot

import (
	"fmt"
)

// SprintTopReposts formats a Top Reposts Bot post and returns the resulting string.
func (s *Stats) SprintTopReposts(limit int) string {
	str := fmt.Sprintf("Top Reposts\n(most reposted messages)\n%s\n", messageDivider)

	topReposts := s.TopReposts(limit)
	for i, repost := range topReposts {
		str += fmt.Sprintf("%d) %d reposts\n", i+1, len(repost.Messages))
		str += fmt.Sprintf("OP: %s (%d)\n", s.Members[repost.OriginalAuthor].Name, repost.AuthorFrequency[repost.OriginalAuthor])

		topRepostersStr := s.SprintTopReposters(limit, repost.Messages[0].Text)
		if topRepostersStr != "" {
			str += fmt.Sprintf("%s\n", topRepostersStr)
		}

		str += fmt.Sprintf("\"%s\"", repost.Messages[0].Text)

		// don't put newline after last ranking
		if i < len(topReposts)-1 {
			str += "\n\n"
		}
	}

	return str
}

// SprintTopReposters formats a Top Reposters Bot post and returns the resulting string.
func (s *Stats) SprintTopReposters(limit int, text string) string {
	repost, ok := s.Reposts[text]
	if !ok {
		repost = &Repost{
			AuthorFrequency: make(map[string]int),
		}
	}

	topReposters := repost.TopReposters(limit)
	if len(topReposters) == 0 {
		return ""
	}

	str := "Reposters: "
	for i, repostAuthor := range topReposters {
		str += fmt.Sprintf("%s (%d)", s.Members[repostAuthor.UserID].Name, repostAuthor.Frequency)

		// don't put newline after last ranking
		if i < len(topReposters)-1 {
			str += ", "
		}
	}

	return str
}
