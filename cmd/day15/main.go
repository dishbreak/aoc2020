package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	rawdata, err := lib.GetInput("inputs/day15.txt")
	if err != nil {
		panic(err)
	}

	input := make([]int, 0)
	for _, part := range strings.Split(rawdata[0], ",") {
		result, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		input = append(input, result)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func playGame(turns int, input []int) int {
	spokenNumbers := make(map[int]int)
	for idx, num := range input {
		spokenNumbers[num] = idx + 1
	}
	var previousNum = input[len(input)-1]
	var currentNum int
	for i := len(input) + 1; i <= turns; i++ {
		lastTurnSpoken, ok := spokenNumbers[previousNum]
		if !ok || lastTurnSpoken == i-1 {
			currentNum = 0
		} else {
			currentNum = i - lastTurnSpoken - 1
		}
		spokenNumbers[previousNum] = i - 1
		previousNum = currentNum
	}

	return previousNum
}

func part1(input []int) int {
	return playGame(2020, input)
}

func part2(input []int) int {
	return playGame(30000000, input)
}
