package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = [][]string{
	[]string{
		`0: 4 1 5`,
		`1: 2 3 | 3 2`,
		`2: 4 4 | 5 5`,
		`3: 4 5 | 5 4`,
		`4: "a"`,
		`5: "b"`,
	},
	[]string{
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	},
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 2, part1(input))
}

func TestParseRule(t *testing.T) {
	assert.Equal(t, rule{number: 0, conditionals: [][]int{
		[]int{4, 1, 5},
	}}, newRule(input[0][0]))
	assert.Equal(t, rule{number: 1, conditionals: [][]int{
		[]int{2, 3},
		[]int{3, 2},
	}}, newRule(input[0][1]))

	assert.Equal(t, rule{number: 4, baseCondition: 'a'}, newRule(input[0][4]))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}
