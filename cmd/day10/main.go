package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := getInput()
	if err != nil {
		panic(err)
	}
	// parts 1 and 2 will rely on sorted data, so we'll sort the input.
	sort.Ints(input)
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []int) int {
	// for part 1, we'll  traverse the sorted array, keeping track of the
	// distribution of differences in an array.
	// note that we're allocating 4 slots just to make it easier to index
	// transitions.
	transitions := make([]int, 4)

	// account for transitions at the beginning and end of the chain.
	transitions[input[0]]++
	transitions[3]++

	for idx, value := range input {
		if idx == 0 {
			continue
		}
		transition := value - input[idx-1]
		transitions[transition]++
	}

	return transitions[1] * transitions[3]
}

func getValueAt(input []int, idx int) int {
	if idx < 0 {
		return 0
	}
	return input[idx]
}

func part2(input []int) int {
	// while it's potentially possible to generate every possible combination of
	// adapters and count the result, we can save a lot of effort by simply
	// counting the number of combinations.

	// to begin, we'll create a state vector with an element for each jolt value
	// in our adapter array.
	// this state vector will represent the number of valid chains that can end
	// with a jolt rating equal to the index
	state := make([]int, input[len(input)-1]+1)

	// next, we'll loop through out input.
	for _, joltValue := range input {
		// if the jolt value is less or equal to 3, it can be a valid chain by itself.
		if joltValue <= 3 {
			state[joltValue] = 1
		}
		// any valid chains ending with an adapter that is rated 1, 2, or 3
		// jolts lower than the current adapter can also end with this adapter,
		// so we'll count their state vector values too.
		state[joltValue] += getValueAt(state, joltValue-1)
		state[joltValue] += getValueAt(state, joltValue-2)
		state[joltValue] += getValueAt(state, joltValue-3)
	}

	// because we know that any useful chain must end with the highest-rated
	// adapter in the set, we can return the state vector entry at the end.
	return state[len(state)-1]
}

func getInput() ([]int, error) {
	f, err := os.Open("inputs/day10.txt")
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		value, err := strconv.Atoi(s.Text())
		if err != nil {
			return result, err
		}
		result = append(result, value)
	}

	return result, nil
}
