package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int{
	0, 3, 6,
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 436, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 175594, part2(input))
}
