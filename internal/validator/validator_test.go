package validator

import (
	"regexp"
	"testing"

	"github.com/lpww/nerdate/internal/assert"
)

func TestValid(t *testing.T) {
	tests := []struct {
		name   string
		errors map[string]string
		want   bool
	}{
		{
			name: "Should return false if there are errors",
			errors: map[string]string{
				"error": "message",
			},
			want: false,
		},
		{
			name:   "Should return true if there are no errors",
			errors: map[string]string{},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			v.Errors = tt.errors
			isValid := v.Valid()

			assert.Equal(t, isValid, tt.want)
		})
	}
}

func TestAddError(t *testing.T) {
	tests := []struct {
		name   string
		errors map[string]string
		want   string
	}{
		{
			name:   "Should add an error",
			errors: map[string]string{},
			want:   "new message",
		},
		{
			name: "Should not override existing errors",
			errors: map[string]string{
				"error": "old message",
			},
			want: "old message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			v.Errors = tt.errors
			v.AddError("error", "new message")

			assert.Equal(t, v.Errors["error"], tt.want)
		})
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		name string
		ok   bool
		want int
	}{
		{
			name: "Should not add an error if the check is ok",
			ok:   true,
			want: 0,
		},
		{
			name: "Should add an error if the check fails",
			ok:   false,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := New()
			v.Check(tt.ok, "error", "this is a message")

			assert.Equal(t, len(v.Errors), tt.want)
		})
	}
}

func TestMatches(t *testing.T) {
	tests := []struct {
		name  string
		value string
		regex regexp.Regexp
		want  bool
	}{
		{
			name:  "Should return true if value matches the regex",
			value: "fake.person@gmail.com",
			regex: *EmailRX,
			want:  true,
		},
		{
			name:  "Should return false if the value does not match the regex",
			value: "test",
			regex: *EmailRX,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches := Matches(tt.value, &tt.regex)

			assert.Equal(t, matches, tt.want)
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		want   bool
	}{
		{
			name:   "Should return true if values are unique",
			values: []string{"hi"},
			want:   true,
		},
		{
			name:   "Should return false if the values are not unique",
			values: []string{"hi", "hi"},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isUnique := Unique(tt.values)

			assert.Equal(t, isUnique, tt.want)
		})
	}
}
