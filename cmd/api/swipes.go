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
		SwipedUserID: input.SwipedUserID,
		Liked:        input.Liked,
	}

	v := validator.New()

	if data.ValidateSwipe(v, swipe); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}
