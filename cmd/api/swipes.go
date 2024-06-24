package main

import (
	"fmt"
	"net/http"
)

func (app *application) createSwipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new swipe")
}
