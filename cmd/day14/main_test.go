package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
	"mem[8] = 11",
	"mem[7] = 101",
	"mem[8] = 0",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, int64(165), part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}

func TestMaskV1(t *testing.T) {
	mask := NewMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	mem := make(map[int64]int64)

	mask.WriteTo(mem, 8, 11)
	assert.Equal(t, int64(73), mem[8])
}
