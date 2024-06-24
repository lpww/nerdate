package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// deferred func will always be run in event of a panic
		defer func() {
			// built in recover() checks if there was a panic
			if err := recover(); err != nil {
				// setting this header triggers go to close the connection after sending a response
				w.Header().Set("Connection", "close")

				// recover() returns an error, so we use Errorf to normalize it into an error and call our server error response helper
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}