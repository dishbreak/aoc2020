package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"1 + 2 * 3 + 4 * 5 + 6",                           // 71
	"1 + (2 * 3) + (4 * (5 + 6))",                     // 51
	"2 * 3 + (4 * 5)",                                 // 26
	"5 + (8 * 3 + 9 + 3 * 4 * 3)",                     // 437
	"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",       // 12240
	"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", // 13632
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 26457, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 694173, part2(input))
}

func TestEvaluatePart2(t *testing.T) {
	assert.Equal(t, 465792, evaluateStatementV2("(5 * 5 * 9 * 6 + (2 + 9 * 5 * 6) + 9) + 7 * 6"))
}
