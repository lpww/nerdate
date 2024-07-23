package data

import (
	"context"
	"database/sql"
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

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	query := `
    INSERT INTO users (name, gender, dob, ascii_art, description, email, password_hash, activated)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id, created_at, version`

	args := []any{user.Name, user.Gender, user.DOB, user.ASCIIArt, user.Description, user.Email, user.Password.hash, user.Activated}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	return nil, nil
}

func (m UserModel) Update(user *User) error {
	return nil
}
