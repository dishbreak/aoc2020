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

	cardLoop := getLoopSize(cardPk)
	doorLoop := getLoopSize(doorPk)

	encKey := getEncKey(cardPk, doorLoop)

	if other := getEncKey(doorPk, cardLoop); encKey != other {
		panic(fmt.Errorf("encryption key mismatch, card got %d, door got %d", other, encKey))
	}

	return encKey
}

func part2(input []int) int {

	return 0
}

func getLoopSize(publicKey int) int {
	// brute-force the loop count.
	// note that for the puzzle input this could take many, many minutes!
	val := 1
	i := 0
	for ; val != publicKey; i++ {

		val = val * 7
		val = val % 20201227
	}
	return i
}

func getEncKey(publicKey, loopSize int) int {
	val := 1
	for i := 1; i <= loopSize; i++ {
		val = val * publicKey
		val = val % 20201227
	}
	return val
}
