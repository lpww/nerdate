package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/lpww/nerdate/internal/validator"
)

type Swipe struct {
	ID string `json:"id"`

	UserID       string `json:"user_id"`
	Liked        bool   `json:"liked"`
	SwipedUserID string `json:"swiped_user_id"`

	CreatedAt time.Time `json:"-"`
	DeletedAt null.Time `json:"-"`
}

func ValidateSwipe(v *validator.Validator, swipe *Swipe) {
	v.Check(swipe.SwipedUserID != "", "swiped_user_id", "must be provided")
	v.Check(validator.IsValidUUID(swipe.SwipedUserID), "swiped_user_id", "must be a valid uuidv4")
}

func (s1 Swipe) Equal(s2 Swipe) bool {
	return s1.UserID == s2.UserID && s1.Liked == s2.Liked && s1.SwipedUserID == s2.SwipedUserID
}

type SwipeModel struct {
	DB *sql.DB
}

func (s SwipeModel) Insert(swipe *Swipe) error {
	query := `
    INSERT INTO swipes (user_id, liked, swiped_user_id)
    VALUES ($1, $2, $3)
    RETURNING id, created_at`

	args := []any{swipe.UserID, swipe.Liked, swipe.SwipedUserID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.DB.QueryRowContext(ctx, query, args...).Scan(&swipe.ID, &swipe.CreatedAt)
}

func (s SwipeModel) Get(id string) (*Swipe, error) {
	return nil, nil
}

func (s SwipeModel) Update(swipe *Swipe) error {
	return nil
}

func (s SwipeModel) Delete(id string) error {
	return nil
}
