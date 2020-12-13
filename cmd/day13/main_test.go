package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"939",
	"7,13,x,x,59,x,31,19",
}

func TestNewBusCalculator(t *testing.T) {
	calc, err := NewBusCalculator(input)
	assert.Nil(t, err)
	assert.Equal(t, &BusCalculator{
		startTime: 939,
		busIds:    []int{7, 13, 59, 31, 19},
		maxId:     59,
	}, calc)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 295, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, part2(input))
}
