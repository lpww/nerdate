package main

import (
	"fmt"
	"net/http"
)

func (app *application) createSwipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		SwipedUserID int64 `json:"swiped_user_id"`
		Liked        bool
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}
