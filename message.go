package groupmestatsbot

import (
	"math"
	"sort"
	"strconv"

	"github.com/MagnusFrater/groupme"
)

func (s *Stats) handleMemberRemovedEvent(event groupme.Event) {
	if event.Type != groupme.MemberRemovedEventType {
		return
	}

	var remover, removed *Member

	for key, i := range event.Data {
		userEventData, ok := groupme.ParseUserEventData(i)
		if ok {
			switch key {
			case groupme.RemoverUserKey:
				remover = NewMember(strconv.Itoa(userEventData.ID), userEventData.Nickname)
			case groupme.RemovedUserKey:
				removed = NewMember(strconv.Itoa(userEventData.ID), userEventData.Nickname)
			}
		}
	}

	s.addKicked(remover, removed)
	s.addKickedBy(remover, removed)
}

func (s *Stats) handleMemberAddedEvent(event groupme.Event) {
	if event.Type != groupme.MemberAddedEventType {
		return
	}

	var adder *Member
	addedUsers := []*Member{}

	for key, value := range event.Data {
		// adder
		if key == groupme.AdderUserKey {
			userEventData, ok := groupme.ParseUserEventData(value)
			if ok {
				adder = NewMember(strconv.Itoa(userEventData.ID), userEventData.Nickname)
				continue
			}
		}

		// added users (array)
		if key == groupme.AddedUsersKey {
			usersEventData, ok := groupme.ParseUsersEventData(value)
			if ok {
				for _, user := range usersEventData {
					addedUsers = append(addedUsers, NewMember(strconv.Itoa(user.ID), user.Nickname))
				}
			}
		}
	}

	s.addAdded(adder, addedUsers)
	s.addAddedBy(adder, addedUsers)
}

// TotalMessages returns the total number of messages.
func (s *Stats) TotalMessages() int {
	return len(s.Messages)
}

// AverageMessageLength returns the average message length.
func (s *Stats) AverageMessageLength() int {
	if s.TotalMessagesLength == 0 || len(s.Messages) == 0 {
		return -1
	}

	return s.TotalMessagesLength / len(s.Messages)
}

// TopMessages returns a sorted list of the most favorited messages.
func (s *Stats) TopMessages(limit int) []*groupme.Message {
	if limit == -1 {
		limit = math.MaxInt64
	}

	sorted := make([]*groupme.Message, len(s.Messages))
	copy(sorted, s.Messages)

	sort.Slice(sorted, func(i, j int) bool { return len(sorted[i].FavoritedBy) > len(sorted[j].FavoritedBy) })

	top := []*groupme.Message{}
	for i := 0; i < limit && i < len(sorted); i++ {
		top = append(top, sorted[i])
	}

	return top
}
