package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := GetInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []int) int {
	complements := make(map[int]int)

	for _, val := range input {
		complement := 2020 - val
		// as we go along, check to see if we've found the complement for this
		// value already. if we have, return the product immediately.
		if _, ok := complements[complement]; ok {
			return val * complement
		}

		// if we haven't found the complement yet, add this value to the set of
		// available complements.
		complements[val] = complement
	}

	return 0
}

func part2(input []int) int {
	// for part 2, complements still are valid, but are semantically a little
	// different. We represent complements as the sum of two values, and their
	// key is their product.
	complements := make(map[int]int)
	// because we need at least 3 items in the slice to get a result, we'll
	// preload the complements map with the first two values.
	complements[input[0]+input[1]] = input[0] * input[1]

	for idx, val := range input {
		// we'll skip the first two entries since we already registered them in
		// the complements map.
		if idx < 2 {
			continue
		}

		// if we have a complmenent that fits the current value, multiply the
		// value against the product and return it.
		complement := 2020 - val
		if product, ok := complements[complement]; ok {
			return val * product
		}

		// if we haven't yet found the matching complement, we'll go back and
		// generate n-1 complements by combining this value with all the
		// previous values in the slice.
		for _, item := range input[:idx] {
			complements[val+item] = val * item
		}
	}

	return 0
}

// GetInput will fetch the input for the problem
func GetInput() ([]int, error) {
	result := make([]int, 0)

	f, err := os.Open("inputs/day1.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, x)
	}
	return result, nil
}
