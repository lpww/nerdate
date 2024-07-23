package data

import (
	"errors"
	"time"

	"github.com/guregu/null"
	"github.com/lpww/nerdate/internal/validator"
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

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not exceed 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not exceed 500 bytes long")

	v.Check(!user.DOB.IsZero(), "dob", "must be provided")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	// this is not a validation becuase it would not be caused by a user error
	// it would only be triggered by a code error, so we panic
	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}
