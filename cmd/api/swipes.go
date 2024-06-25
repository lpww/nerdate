package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) createSwipeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		SwipedUserID int64 `json:"swiped_user_id"`
		Liked        bool
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}
