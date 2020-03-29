package groupmestatsbot

import (
	"fmt"
	"sort"

	"github.com/MagnusFrater/groupme"
)

// Repost is a counter of message duplicates.
type Repost struct {
	Messages        []*groupme.Message
	AuthorFrequency map[string]int // UserID -> frequency

	OriginalAuthor    string // UserID of the first author of the message
	OriginalCreatedAt int    // timestamp of the first message (integer seconds since the UNIX epoch)
}

// RepostAuthor is a counter of authors of a repost.
type RepostAuthor struct {
	UserID    string
	Frequency int
}

func (r *Repost) addMessage(message *groupme.Message) {
	if message.Text == "" {
		return
	}

	r.Messages = append(r.Messages, message)
}

func (r *Repost) addAuthor(userID string) {
	if _, ok := r.AuthorFrequency[userID]; !ok {
		r.AuthorFrequency[userID] = 0
	}
}

func (r *Repost) incAuthor(userID string) {
	r.addAuthor(userID)
	r.AuthorFrequency[userID]++
}

func (r *Repost) updateOriginalAuthor(message *groupme.Message) {
	if r.OriginalCreatedAt == 0 {
		r.OriginalAuthor = message.UserID
		r.OriginalCreatedAt = message.CreatedAt
		return
	}

	if message.CreatedAt < r.OriginalCreatedAt {
		r.OriginalAuthor = message.UserID
		r.OriginalCreatedAt = message.CreatedAt
	}
}

func (s *Stats) addRepost(message *groupme.Message) {
	if message.Text == "" {
		return
	}

	if repost, ok := s.Reposts[message.Text]; !ok {
		s.Reposts[message.Text] = &Repost{
			Messages:        []*groupme.Message{message},
			AuthorFrequency: make(map[string]int),
		}
	} else {
		repost.addMessage(message)
	}
}

func (s *Stats) incRepost(message *groupme.Message) {
	if message.Text == "" {
		return
	}

	s.addRepost(message)

	s.Reposts[message.Text].incAuthor(message.UserID)
	s.Reposts[message.Text].updateOriginalAuthor(message)
}

// TopReposts returns a sorted list of the most duplicated messages.
func (s *Stats) TopReposts(limit int) []*Repost {
	sorted := []*Repost{}
	for _, repost := range s.Reposts {
		sorted = append(sorted, repost)
	}

	sort.Slice(sorted, func(i, j int) bool {
		if len(sorted[i].TopReposters(limit)) == 0 {
			return false
		}

		return len(sorted[i].Messages) > len(sorted[j].Messages)
	})

	top := []*Repost{}
	for i := 0; i < limit && i < len(sorted); i++ {
		// don't track messages with zero reposters
		if len(sorted[i].TopReposters(limit)) == 0 {
			continue
		}

		top = append(top, sorted[i])
	}

	return top
}

// TopReposters returns a sorted list of who reposted this message the most.
func (r *Repost) TopReposters(limit int) []*RepostAuthor {
	sorted := []*RepostAuthor{}
	for userID, frequency := range r.AuthorFrequency {
		// original author is not a reposter
		if userID == r.OriginalAuthor {
			continue
		}

		sorted = append(sorted, &RepostAuthor{
			UserID:    userID,
			Frequency: frequency,
		})
	}

	// no reposters
	if len(sorted) == 0 {
		return []*RepostAuthor{}
	}

	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Frequency > sorted[j].Frequency })

	top := []*RepostAuthor{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

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