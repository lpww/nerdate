package main

import (
	"bytes"
	"encoding/json"
	"net/http"
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
				UserID:       "uuid-100",
				Liked:        true,
				SwipedUserID: "uuid-101",
			},
		},
		{
			name: "Partial swipe",
			give: data.Swipe{
				UserID:       "uuid-100",
				SwipedUserID: "uuid-101",
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
				ID           string `json:"id"`
				UserID       string `json:"user_id"`
				Liked        bool   `json:"liked"`
				SwipedUserID string `json:"swiped_user_id"`
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

func TestWriteJSON(t *testing.T) {
	app := newTestApplication()

	tests := []struct {
		name string
		give envelope
	}{
		{
			name: "Standard swipe",
			give: envelope{
				"swipe": data.Swipe{
					UserID:       "uuid-100",
					Liked:        true,
					SwipedUserID: "uuid-101",
				}},
		},
		{
			name: "Partial swipe",
			give: envelope{
				"swipe": data.Swipe{
					UserID:       "uuid-100",
					SwipedUserID: "uuid-101",
				}},
		},
		{
			name: "Empty swipe",
			give: envelope{
				"swipe": data.Swipe{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			if err := app.writeJSON(rr, http.StatusOK, tt.give, nil); err != nil {
				t.Fatalf("writeJSON returned an unexpected error: %v", err)
			}

			if rr.Code != http.StatusOK {
				t.Errorf("unexpected status code: got %d, want %d", rr.Code, http.StatusOK)
			}

			if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
				t.Errorf("unexpected content type: got %s, want application/json", contentType)
			}

			var responseBody envelope
			if err := json.Unmarshal(rr.Body.Bytes(), &responseBody); err != nil {
				t.Fatalf("failed to unmarshal response JSON: %v", err)
			}

			if _, ok := responseBody["swipe"]; !ok {
				t.Fatal("no swipe key in response body")
			}

			var movieMap map[string]interface{}
			var ok bool
			if movieMap, ok = responseBody["swipe"].(map[string]interface{}); !ok {
				t.Fatal("response body swipe is not in the map")
			}

			var input struct {
				UserID       string `json:"user_id"`
				Liked        bool   `json:"liked"`
				SwipedUserID string `json:"swiped_user_id"`
			}

			js, err := json.Marshal(movieMap)
			if err != nil {
				t.Fatal("error while marshalling swipe map")
			}
			json.Unmarshal(js, &input)

			swipe := data.Swipe{
				UserID:       input.UserID,
				Liked:        input.Liked,
				SwipedUserID: input.SwipedUserID,
			}

			if _, ok := tt.give["swipe"]; !ok {
				panic("no swipe key in literal envelope")
			}

			if !swipe.Equal(tt.give["swipe"].(data.Swipe)) {
				t.Errorf("got: %+v; want: %+v", swipe, tt.give["swipe"])
			}
		})
	}
}
