package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func SeatIDFromPass(pass string) int {
	rowSymbols := string(pass[:7])
	colSymbols := string(pass[7:len(pass)])
	row := binarySearch(rowSymbols, 'F', 'B', 0, 127)
	col := binarySearch(colSymbols, 'L', 'R', 0, 7)

	return (row * 8) + col
}

func binarySearch(input string, low, high rune, min, max int) int {
	for _, x := range input {
		switch x {
		case high:
			min = (min+max)/2 + 1
		case low:
			max = (min + max) / 2
		}
	}
	return min
}

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	highestPassID := SeatIDFromPass(input[0])

	for idx, val := range input {
		if idx == 0 {
			continue
		}

		newPassID := SeatIDFromPass(val)
		if newPassID > highestPassID {
			highestPassID = newPassID
		} else {
			break
		}
	}

	return highestPassID
}

func part2(input []string) int {
	previousPass := input[0]

	for idx, val := range input {
		if idx == 0 {
			continue
		}

		if previousPass[9] != val[9] {
			previousPass = val
		} else {
			break
		}
	}
	return SeatIDFromPass(previousPass) + 1
}

func getInput() ([]string, error) {
	f, err := os.Open("inputs/day5.txt")
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	sort.Strings(result)
	return result, nil
}
