package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Swipes SwipeModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Swipes: SwipeModel{DB: db},
	}
}
