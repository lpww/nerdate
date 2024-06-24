package data

import "time"

type Swipe struct {
	ID int64 `json:"id"`

	UserID       int64 `json:"user_id"` // todo: use the user.ID field when it's been created
	Liked        bool  `json:"liked"`
	SwipedUserID int64 `json:"swiped_user_id"` // todo: use the user.ID field when it's been created

	CreatedAt time.Time `json:"-"`
}

func (s1 Swipe) Equal(s2 Swipe) bool {
	return s1.UserID == s2.UserID && s1.Liked == s2.Liked && s1.SwipedUserID == s2.SwipedUserID
}
