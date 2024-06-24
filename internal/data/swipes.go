package data

import "time"

type Swipe struct {
	ID int64

	UserID       int64 // todo: use the user.ID field when it's been created
	Liked        bool
	SwipedUserID int64 // todo: use the user.ID field when it's been created

	CreatedAt time.Time
}
