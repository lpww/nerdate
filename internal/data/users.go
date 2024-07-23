package data

import (
	"errors"
	"time"

	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   null.Time `json:"-"`
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	DOB         time.Time `json:"dob"`
	ASCIIArt    string    `json:"ascii_art"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	Password    password  `json:"-"`
	Activated   bool      `json:"activated"`
	Version     int       `json:"-"`
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
