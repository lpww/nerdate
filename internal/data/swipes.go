package data

import (
	"time"

	"github.com/lpww/nerdate/internal/validator"
)

type Swipe struct {
	ID string `json:"id"`

	UserID       string `json:"user_id"` // todo: use the user.ID field when it's been created
	Liked        bool   `json:"liked"`
	SwipedUserID string `json:"swiped_user_id"` // todo: use the user.ID field when it's been created

	CreatedAt time.Time `json:"-"`
}

func ValidateSwipe(v *validator.Validator, swipe *Swipe) {
	v.Check(swipe.SwipedUserID != "", "swiped_user_id", "must be provided")
}

func (s1 Swipe) Equal(s2 Swipe) bool {
	return s1.UserID == s2.UserID && s1.Liked == s2.Liked && s1.SwipedUserID == s2.SwipedUserID
}
