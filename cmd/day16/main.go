package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day16.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

var rangeMatcher = regexp.MustCompile(`^([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)$`)

func part1(input []string) int {
	counter := 0

	intervals := make([]*lib.Range, 0)
	for ; input[counter] != ""; counter++ {
		results := rangeMatcher.FindAllStringSubmatch(input[counter], -1)
		if len(results) == 1 {
			name := results[0][1]
			bounds := make([]int, 4)
			for idx, value := range results[0][2:] {
				parsed, _ := strconv.Atoi(value)
				bounds[idx] = parsed
			}
			intervals = append(
				intervals,
				&lib.Range{Min: bounds[0], Max: bounds[1], Metadata: name},
				&lib.Range{Min: bounds[2], Max: bounds[3], Metadata: name},
			)
		}
	}

	tree, err := lib.NewIntervalTree(intervals)
	if err != nil {
		panic(err)
	}

	counter += 5

	errorRate := 0
	for ; counter < len(input); counter++ {
		if input[counter] == "" {
			continue
		}
		parts := strings.Split(input[counter], ",")
		for _, fieldValue := range parts {
			parsed, err := strconv.Atoi(fieldValue)
			if err != nil {
				panic(err)
			}
			if len(tree.Find(parsed)) == 0 {
				for _, interval := range intervals {
					if interval.Contains(parsed) {
						panic(fmt.Errorf("%d is in range %v", parsed, interval))
					}
				}
				errorRate += parsed
			}
		}
	}

	return errorRate
}

func part2(input []string) int {
	return 0
}
