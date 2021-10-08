package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day22.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input [][]string) int {
	game := buildGame(input)

	for game.playRound() {
	}

	return game.scoreGame()
}

func part2(input [][]string) int {
	return 0
}
