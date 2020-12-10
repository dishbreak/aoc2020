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
	fmt.Printf("Part 2: %d\n", part2(input, 25))
}

func part1(input []int) int {
	result, _ := findWeakness(input, 25)

	return result
}

func findWeakness(input []int, preamble int) (int, int) {
	// we'll use an array to keep track of the previous 25 numbers in the
	// sequence.
	// Start by loading the array with the 25-number preamble
	preceding := make([]int, preamble)
	copy(preceding, input)

	for i := preamble; i < len(input); i++ {
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
		preceding[i%preamble] = input[i]
	}

	return -1, -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part2(input []int, preamble int) int {
	// we're going to use a little dynamic programming here to build on top of
	// our answer from part 1.

	// first, find the actual weakness and its index
	weakness, _ := findWeakness(input, preamble)

	// next, create state vectors for the sum, min, and max values
	sums := make([]int, len(input))
	mins := make([]int, len(input))
	maxs := make([]int, len(input))
	copy(sums, input)
	copy(mins, input)
	copy(maxs, input)

	// we're going to try window sizes ranging from 2 to the entire input.
	for windowSize := 2; windowSize <= len(input); windowSize++ {
		// this is the fun part. by starting at the end of the state vectors and
		// moving backwards, we can actually build up our windowed sums, mins,
		// and maxes!
		// this works because the sum of 3 contiguous integers ending at index 5
		// will be the sum of the value at index 5 and the sum of the 2
		// contiguous integers ending at index 4.
		// by working backwards, we can keep the same state vector for all iterations.
		for i := len(input) - 1; i >= windowSize-1; i-- {
			mins[i] = min(input[i], mins[i-1])
			maxs[i] = max(input[i], maxs[i-1])
			sums[i] = input[i] + sums[i-1]
			if sums[i] == weakness {
				return mins[i] + maxs[i]
			}
		}
	}
	return 0
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
