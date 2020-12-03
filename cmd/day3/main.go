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
}

func part1(input []string) int {
	return traverse(input, 1, 3)
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
