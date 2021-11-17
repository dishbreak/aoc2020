package main

import (
	"fmt"
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	rawInput, err := lib.GetInput("inputs/day25.txt")
	if err != nil {
		panic(err)
	}

	input := make([]int, 2)
	for idx, _ := range input {
		parsed, _ := strconv.Atoi(rawInput[idx])
		input[idx] = parsed
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []int) int {
	doorPk := input[0]
	cardPk := input[1]

	loopSize, err := getLoopSize(doorPk)
	if err != nil {
		panic(err)
	}

	return getEncKey(cardPk, loopSize)
}

func part2(input []int) int {

	return 0
}

func getLoopSize(publicKey int) (int, error) {
	for i := 1; i <= 20; i++ {
		val := 1
		for j := 0; j < i; j++ {
			val = val * 7
			val = val % 20201227
		}
		if val == publicKey {
			return i, nil
		}
	}
	return -1, fmt.Errorf("failed to break key %d", publicKey)
}

func getEncKey(publicKey, loopSize int) int {
	encKey := publicKey
	for i := 0; i < loopSize; i++ {
		encKey = encKey * 7
		encKey = encKey % 20201227
	}
	return encKey
}
