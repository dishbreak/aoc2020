package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"L.LL.LL.LL",
	"LLLLLLL.LL",
	"L.L.L..L..",
	"LLLL.LL.LL",
	"L.LL.LL.LL",
	"L.LLLLL.LL",
	"..L.L.....",
	"LLLLLLLLLL",
	"L.LLLLLL.L",
	"L.LLLLL.LL",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 37, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 26, part2(input))
}
