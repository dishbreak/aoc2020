package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = "389125467"

func TestPart1(t *testing.T) {
	assert.Equal(t, "67384529", part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}
