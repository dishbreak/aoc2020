package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []int) int {
	result, _ := findWeakness(input)

	return result
}

func findWeakness(input []int) (int, int) {
	var preceding [25]int

	for i := 0; i < 25; i++ {
		preceding[i] = input[i]
	}

	for i := 25; i < len(input); i++ {
		matchFound := false
		for j := 0; j < len(preceding); j++ {
			if preceding[j] > input[i] {
				continue
			}
			for k := 0; k < len(preceding); k++ {
				if j == k {
					continue
				}
				if preceding[j]+preceding[k] == input[i] {
					matchFound = true
					break
				}
			}
			if matchFound {
				break
			}
		}

		if !matchFound {
			return input[i], i
		}
		preceding[i%25] = input[i]
	}

	return -1, -1
}

func getInput() ([]int, error) {
	f, err := os.Open("inputs/day9.txt")
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)

	s := bufio.NewScanner(f)
	for s.Scan() {
		item, err := strconv.Atoi(s.Text())
		if err != nil {
			return result, err
		}

		result = append(result, item)
	}

	return result, nil
}
