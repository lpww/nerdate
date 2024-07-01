package main

import (
	"fmt"
	"net/http"

	"github.com/lpww/nerdate/internal/data"
	"github.com/lpww/nerdate/internal/validator"
)

func (app *application) createSwipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		SwipedUserID string `json:"swiped_user_id"`
		Liked        bool   `json:"liked"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	swipe := &data.Swipe{
		UserID:       "51824827-85f8-4e32-854b-5d10da52446a", // todo: make this dynamic when auth has been implemented
		SwipedUserID: input.SwipedUserID,
		Liked:        input.Liked,
	}

	v := validator.New()

	if data.ValidateSwipe(v, swipe); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Swipes.Insert(swipe)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// add a location header so the user knows which url they can find the resource in
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/swipes/%s", swipe.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"swipe": swipe}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
