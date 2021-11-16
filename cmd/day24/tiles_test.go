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
