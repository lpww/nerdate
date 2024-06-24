package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/lpww/nerdate/internal/data"
)

func TestReadJSON(t *testing.T) {
	app := newTestApplication()

	tests := []struct {
		name string
		give data.Swipe
	}{
		{
			name: "Standard swipe",
			give: data.Swipe{
				UserID:       100,
				Liked:        true,
				SwipedUserID: 101,
			},
		},
		{
			name: "Partial swipe",
			give: data.Swipe{
				UserID:       100,
				SwipedUserID: 101,
			},
		},
		{
			name: "Empty swipe",
			give: data.Swipe{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js, err := json.Marshal(tt.give)
			if err != nil {
				t.Fatalf("failed to marshal JSON: %v", err)
			}

			var input struct {
				ID           int64 `json:"id"`
				UserID       int64 `json:"user_id"`
				Liked        bool  `json:"liked"`
				SwipedUserID int64 `json:"swiped_user_id"`
			}

			writeRecorder := httptest.NewRecorder()
			readRequest := httptest.NewRequest("POST", "/v1/swipes/", bytes.NewBuffer(js))

			err = app.readJSON(writeRecorder, readRequest, &input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			readSwipe := data.Swipe{
				UserID:       input.UserID,
				Liked:        input.Liked,
				SwipedUserID: input.SwipedUserID,
			}

			if !tt.give.Equal(readSwipe) {
				t.Errorf("got: %v; want %v", readSwipe, tt.give)
			}
		})
	}
}
