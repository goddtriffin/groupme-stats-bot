package groupmestatsbot

import (
	"math"
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
	if limit == -1 {
		limit = math.MaxInt64
	}

	sorted := []*Repost{}
	for _, repost := range s.Reposts {
		sorted = append(sorted, repost)
	}

	sort.Slice(sorted, func(i, j int) bool { return len(sorted[i].TopReposters(-1)) > len(sorted[j].TopReposters(-1)) })

	top := []*Repost{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}

// TopReposters returns a sorted list of who reposted this message the most.
func (r *Repost) TopReposters(limit int) []*RepostAuthor {
	if limit == -1 {
		limit = math.MaxInt64
	}

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
