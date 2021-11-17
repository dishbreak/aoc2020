package main

import (
	"testing"

	"github.com/dishbreak/aoc2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	input := "seswneswswsenwwnwse"
	expected := []direction{se, sw, ne, sw, sw, se, nw, w, nw, se}

	assert.Equal(t, expected, tokenize(input))
}

func TestTraverse(t *testing.T) {
	input := []direction{se, sw, ne, sw, sw, se, nw, w, nw, se}
	expected := lib.Point3D{X: -3, Y: 3, Z: 0}
	assert.Equal(t, expected, traverse(input))
}

func TestNeighbors(t *testing.T) {
	type testCase struct {
		tiles    map[lib.Point3D]int
		expected int
	}

	cases := map[string]testCase{
		"3 neighbors found": {
			tiles: map[lib.Point3D]int{
				{X: 0, Y: -1, Z: 1}: 1,
				{X: 0, Y: 1, Z: -1}: 1,
				{X: 1, Y: -1, Z: 0}: 1,
			},
			expected: 3,
		},
		"no neighbors found": {
			tiles: map[lib.Point3D]int{
				{X: 0, Y: -5, Z: 5}: 1,
				{X: 0, Y: 2, Z: -2}: 1,
			},
			expected: 0,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			l := lobbyFloor{
				tiles: tc.tiles,
			}
			assert.Equal(t, tc.expected, l.countNeighbors(lib.Point3D{}))
		})
	}

}
