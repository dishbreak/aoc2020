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
	col, count := 0, 0
	for _, row := range input {
		switch row[col] {
		case '#':
			count++
		}
		col = (col + 3) % len(row)
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
