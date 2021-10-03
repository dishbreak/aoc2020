package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntesection(t *testing.T) {
	type testCase struct {
		one      []string
		other    []string
		expected []string
	}

	testCases := map[string]testCase{
		"happy path": {
			[]string{"apple", "banana", "cantelope", "daikon"},
			[]string{"banana", "daikon", "awtermelon"},
			[]string{"banana", "daikon"},
		},
		"full intersection": {
			[]string{"apple", "banana", "cantelope", "daikon"},
			[]string{"apple", "banana", "cantelope", "daikon"},
			[]string{"apple", "banana", "cantelope", "daikon"},
		},
		"no intesection": {
			[]string{"strawberry", "mango", "cantelope", "daikon"},
			[]string{"banana", "kiwi", "awtermelon"},
			[]string{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, intersection(tc.one, tc.other))
		})
	}
}
