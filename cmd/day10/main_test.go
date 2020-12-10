package main

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestPart1(t *testing.T) {
	input := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	sort.Ints(input)
	assert.Equal(t, 220, part1(input))
}

func TestPart2(t *testing.T) {
	input := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	sort.Ints(input)
	assert.Equal(t, 19208, part2(input))
}
