package internal_test

import (
	"testing"

	"github.com/selesy/golang_enumerations/internal"
	"gotest.tools/assert"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name string
		inp  string
	}{
		{"Mixed", "NormalizeMe"},
		{"Upper", "NORMALIZEME"},
		{"Lower", "normalizeme"},
		{"Special", "^Normalize_Me$"},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, "normalizeme", internal.Normalize(test.inp))
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name string
		inp  []string
		exp  bool
	}{
		{"Unique", []string{"one", "two", "three", "four", "five"}, true},
		{"Repeats", []string{"one", "two", "three", "four", "four"}, false},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.exp, internal.Unique(test.inp))
		})
	}
}
