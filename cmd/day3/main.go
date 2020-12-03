package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	return traverse(input, 1, 3)
}

func part2(input []string) int {
	result := 1

	slopes := [][]int{
		[]int{1, 1},
		[]int{1, 3},
		[]int{1, 5},
		[]int{1, 7},
		[]int{2, 1},
	}

	for _, slope := range slopes {
		result = result * traverse(input, slope[0], slope[1])
	}

	return result
}

func traverse(input []string, down, right int) int {
	col, count := 0, 0
	for idx, row := range input {
		if idx%down != 0 {
			continue
		}
		switch row[col] {
		case '#':
			count++
		}
		col = (col + right) % len(row)
	}

	return count
}

func getInput() ([]string, error) {
	result := make([]string, 0)

	f, err := os.Open("inputs/day3.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}
