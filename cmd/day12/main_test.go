package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"F10",
	"N3",
	"F7",
	"R90",
	"F11",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 25, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 286, part2(input))
}
