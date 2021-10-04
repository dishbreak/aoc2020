package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
	"trh fvjkl sbzzf mxmxvkd (contains dairy)",
	"sqjhc fvjkl (contains soy)",
	"sqjhc mxmxvkd sbzzf (contains fish)",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}
