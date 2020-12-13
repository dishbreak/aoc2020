package main

import (
	"fmt"
	"image"
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day12.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

var headingVectors = []image.Point{
	image.Point{0, 1},
	image.Point{1, 0},
	image.Point{0, -1},
	image.Point{-1, 0},
}

type Ship struct {
	position image.Point
	heading  int
}

func (s *Ship) GiveInstruction(instruction string) {
	if len(instruction) == 0 {
		return
	}

	argument, err := strconv.Atoi(string(instruction[1:]))
	if err != nil {
		panic(err)
	}

	switch instruction[0] {
	case 'N':
		vector := headingVectors[0].Mul(argument)
		s.position = s.position.Add(vector)
	case 'E':
		s.position = s.position.Add(headingVectors[1].Mul(argument))
	case 'S':
		s.position = s.position.Add(headingVectors[2].Mul(argument))
	case 'W':
		s.position = s.position.Add(headingVectors[3].Mul(argument))
	case 'R':
		s.heading += argument
	case 'L':
		s.heading -= argument
	case 'F':
		vectorIdx := (s.heading%360/90 + 4) % 4
		s.position = s.position.Add(headingVectors[vectorIdx].Mul(argument))
	}
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func (s *Ship) ManhattanDist() int {
	return abs(s.position.X) + abs(s.position.Y)
}

func part1(input []string) int {
	s := &Ship{
		heading: 90,
	}

	for _, instruction := range input {
		s.GiveInstruction(instruction)
	}

	return s.ManhattanDist()
}
