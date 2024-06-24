package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	// "net/http/httptest"
	//
	// "github.com/lpww/govs/app"
)

func TestHealthcheckHandler(t *testing.T) {
	app := newTestApplication()
	ts := newTestServer(app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/v1/healthcheck")
	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	expected := `{
    "status": "available",
    "system_info": {
      "environment": "testing",
      "version": "1.0.0"
    }
  }
  `

	expected = strings.ReplaceAll(expected, " ", "")
	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, "\n", "")

	got := string(body)
	got = strings.ReplaceAll(got, " ", "")
	got = strings.ReplaceAll(got, "\t", "")
	got = strings.ReplaceAll(got, "\n", "")

	assert.Equal(t, got, expected)
}
