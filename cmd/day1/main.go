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
		if _, ok := complements[complement]; ok {
			return val * complement
		}
		complements[val] = complement
	}

	return 0
}

func part2(input []int) int {
	complements := make(map[int]int)
	complements[input[0]+input[1]] = input[0] * input[1]

	for idx, val := range input {
		if idx < 2 {
			continue
		}

		complement := 2020 - val
		if product, ok := complements[complement]; ok {
			return val * product
		}

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
