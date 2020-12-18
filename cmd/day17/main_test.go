package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	".#.",
	"..#",
	"###",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 112, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 848, part2(input))
}
