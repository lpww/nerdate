package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/lpww/nerdate/internal/data"
	"github.com/lpww/nerdate/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string    `json:"name"`
		Gender      string    `json:"gender"`
		DOB         time.Time `json:"dob"`
		ASCIIArt    string    `json:"ascii_art"`
		Description string    `json:"description"`
		Email       string    `json:"email"`
		Password    string    `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:        input.Name,
		Gender:      input.Gender,
		DOB:         input.DOB,
		ASCIIArt:    input.ASCIIArt,
		Description: input.Description,
		Email:       input.Email,
		Activated:   true, // todo: set this to false once email confirmation is implemtented
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		// if duplicate email found, return a validation error
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) discoverUsersHandler(w http.ResponseWriter, r *http.Request) {
	// todo: add pagination, filters, and sorting based on query string
	// todo: validate query string input
	// todo: ignore people already swiped on
	// todo: don't return the logged in user to themselves
	// todo: filter options should be applied based on the logged in users search preferences
	users, err := app.models.Users.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"users": users}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
