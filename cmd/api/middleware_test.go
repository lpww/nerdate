package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lpww/nerdate/internal/assert"
)

func TestRateLimit(t *testing.T) {
	maxBurst := 4
	tReqs := 6

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	app := newTestApplication()
	limited := app.rateLimit(next)

	results := []*http.Response{}

	// repeat 5 times
	for i := 0; i < tReqs; i++ {
		rr := httptest.NewRecorder()

		r, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		// normally set by HTTP server automatically before invoking a handler
		// the test calls the middleware directly, bypassing the server call
		// so we need to set the value manually
		r.RemoteAddr = "1.1.1.1:3000"

		limited.ServeHTTP(rr, r)

		rs := rr.Result()
		results = append(results, rs)
	}

	for i := 0; i < maxBurst; i++ {
		rs := results[i]
		assert.Equal(t, rs.StatusCode, http.StatusOK)

		defer rs.Body.Close()
		body, err := io.ReadAll(rs.Body)
		if err != nil {
			t.Fatal(err)
		}
		body = bytes.TrimSpace(body)

		assert.Equal(t, string(body), "OK")
	}

	expected := `{
    "error": "rate limit exceeded"
  }
  `

	expected = strings.ReplaceAll(expected, " ", "")
	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, "\n", "")

	for i := maxBurst; i < tReqs; i++ {
		rs := results[i]
		assert.Equal(t, rs.StatusCode, http.StatusTooManyRequests)

		defer rs.Body.Close()
		body, err := io.ReadAll(rs.Body)
		if err != nil {
			t.Fatal(err)
		}
		body = bytes.TrimSpace(body)

		got := string(body)
		got = strings.ReplaceAll(got, " ", "")
		got = strings.ReplaceAll(got, "\t", "")
		got = strings.ReplaceAll(got, "\n", "")

		assert.Equal(t, got, expected)
	}
}
