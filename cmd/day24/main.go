package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day24.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	l := &lobbyFloor{}
	l.tiles = make(map[lib.Point3D]int)

	for _, step := range input {
		l.flipTile(step)
	}

	return l.countBlackTiles()
}

func part2(input []string) int {
	return 0
}
