package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = [][]string{
	{
		"Player 1:",
		"9",
		"2",
		"6",
		"3",
		"1",
	},
	{
		"Player 2:",
		"5",
		"8",
		"4",
		"7",
		"10",
	},
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 306, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 291, part2(input))
}
