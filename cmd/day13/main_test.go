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
		maxID:     59,
	}, calc)
}

func TestLoadSchedule(t *testing.T) {
	expected := []BusScheduleEntry{
		BusScheduleEntry{59, 4},
		BusScheduleEntry{31, 6},
		BusScheduleEntry{19, 7},
		BusScheduleEntry{13, 1},
		BusScheduleEntry{7, 0},
	}
	actual := LoadSchedule(input[1])
	assert.Equal(t, expected, actual)
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 295, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 1068781, part2(input))
}
