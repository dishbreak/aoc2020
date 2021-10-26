package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() *shellGame {
	return &shellGame{
		cups: []int{3, 7, 8, 9, 5},
	}
}
func TestRemoveAt(t *testing.T) {

	type testCase struct {
		indexToRemove int
		val           int
		result        []int
	}

	testCases := map[string]testCase{
		"no wraparound": {
			indexToRemove: 3,
			val:           9,
			result:        []int{3, 7, 8, 5},
		},
		"removing start": {
			indexToRemove: 0,
			val:           3,
			result:        []int{7, 8, 9, 5},
		},
		"removing start wraparound": {
			indexToRemove: 5,
			val:           3,
			result:        []int{7, 8, 9, 5},
		},
		"removing end": {
			indexToRemove: 4,
			val:           5,
			result:        []int{3, 7, 8, 9},
		},
		"removing end wraparound": {
			indexToRemove: 9,
			val:           5,
			result:        []int{3, 7, 8, 9},
		},
		"negative wraparound": {
			indexToRemove: -2,
			val:           9,
			result:        []int{3, 7, 8, 5},
		},
		"positive wraparound": {
			indexToRemove: 6,
			val:           7,
			result:        []int{3, 8, 9, 5},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			g := setup()
			val := g.removeAt(tc.indexToRemove)
			assert.Equal(t, tc.val, val)
			assert.Equal(t, tc.result, g.cups)
		})
	}
}

func TestInsertAt(t *testing.T) {
	type testCase struct {
		index  int
		value  int
		result []int
	}

	testCases := map[string]testCase{
		"insert middle": {
			index:  2,
			value:  2,
			result: []int{3, 7, 2, 8, 9, 5},
		},
		"insert wraparound": {
			index:  7,
			value:  2,
			result: []int{3, 7, 2, 8, 9, 5},
		},
		"insert start": {
			index:  0,
			value:  2,
			result: []int{2, 3, 7, 8, 9, 5},
		},
		"insert end": {
			index:  4,
			value:  2,
			result: []int{3, 7, 8, 9, 2, 5},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			g := setup()
			g.insertAt(tc.value, tc.index)
			assert.Equal(t, tc.result, g.cups)
		})
	}
}
