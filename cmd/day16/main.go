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

func toIntSlice(input []string) ([]int, error) {
	output := make([]int, len(input))
	for idx, value := range input {
		parsed, err := strconv.Atoi(value)
		if err != nil {
			return output, err
		}
		output[idx] = parsed
	}
	return output, nil
}

func validTicket(fieldVals []int, tree *lib.IntervalTreeNode) bool {
	for _, fieldVal := range fieldVals {
		if len(tree.Find(fieldVal)) == 0 {
			return false
		}
	}
	return true
}

func parseRanges(input []string) ([]*lib.Range, int) {
	counter := 0

	intervals := make([]*lib.Range, 0)
	for ; input[counter] != ""; counter++ {
		results := rangeMatcher.FindAllStringSubmatch(input[counter], -1)
		if len(results) == 1 {
			name := results[0][1]
			bounds, _ := toIntSlice(results[0][2:])
			intervals = append(
				intervals,
				&lib.Range{Min: bounds[0], Max: bounds[1], Metadata: name},
				&lib.Range{Min: bounds[2], Max: bounds[3], Metadata: name},
			)
		}
	}

	return intervals, counter
}

var rangeMatcher = regexp.MustCompile(`^([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)$`)

func part1(input []string) int {

	intervals, counter := parseRanges(input)

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
		fieldVals, err := toIntSlice(strings.Split(input[counter], ","))
		if err != nil {
			panic(err)
		}

		for _, fieldVal := range fieldVals {
			if len(tree.Find(fieldVal)) == 0 {
				errorRate += fieldVal
			}
		}
	}

	return errorRate
}

func populateTicketFields(myTicket []int, validTickets [][]int, tree *lib.IntervalTreeNode) map[string]int {
	myPopulatedTicket := make(map[string]int, len(myTicket))

	// for each position, we're going to keep track of how many fields match the
	// value in that position, using our interval tree.
	fieldsVector := make([]map[string]int, len(myTicket))
	for idx := range fieldsVector {
		fieldsVector[idx] = make(map[string]int)
	}

	// loop through all the valid tickets.
	for _, ticket := range validTickets {
		// for each value, find the matching intervals and increment the counter
		// for the given field name
		for idx, fieldVal := range ticket {
			matchingIntervals := tree.Find(fieldVal)
			for j := range matchingIntervals {
				if fieldName, ok := matchingIntervals[j].Metadata.(string); ok {
					fieldsVector[idx][fieldName]++
				}
			}
		}
	}

	// we're going to play multiple rounds of the following game:
	// 1. for each position, see if only 1 field matches all values in that
	// position.
	// 2. If we find a winner, write the value from that position on our ticket
	// into the map, using the field name as the key.
	// 3. eliminate the field from all our field vectors and eliminate the
	// position from consideration.

	// We'll keep playing rounds until the map has entries for each position in
	// the ticket.
	for len(myPopulatedTicket) != len(myTicket) {
		// find fields that match all values in a given position
		for idx, state := range fieldsVector {
			// if a state vector entry is nil it means we've locked that
			// position to a given field.
			if state == nil {
				continue
			}
			var fieldCandidate string
			matches := 0
			for field, hits := range state {
				if hits == len(validTickets) {
					fieldCandidate = field
					matches++
				}
			}
			// if only one field matched all values for this ticket, we can lock
			// in this position as matching to that field
			if matches == 1 {
				myPopulatedTicket[fieldCandidate] = myTicket[idx]
				fieldsVector[idx] = nil
			}
		}

		// eliminate locked in fields from the remaining state vectors
		for _, state := range fieldsVector {
			if state == nil {
				continue
			}
			// ensure all fields that we've solved aren't present in the vector.
			// this will let us find matches in successive rounds.
			for k := range myPopulatedTicket {
				delete(state, k)
			}
		}

	}

	return myPopulatedTicket
}

func part2(input []string) int {
	intervals, counter := parseRanges(input)

	tree, err := lib.NewIntervalTree(intervals)
	if err != nil {
		panic(err)
	}

	counter += 2

	myTicketValues, err := toIntSlice(strings.Split(input[counter], ","))
	if err != nil {
		panic(err)
	}

	counter += 3
	validTickets := make([][]int, 0)

	for ; counter < len(input); counter++ {
		if input[counter] == "" {
			continue
		}

		ticket, err := toIntSlice(strings.Split(input[counter], ","))
		if err != nil {
			panic(err)
		}
		if validTicket(ticket, tree) {
			validTickets = append(validTickets, ticket)
		}
	}

	myTicketFields := populateTicketFields(myTicketValues, validTickets, tree)

	result := 1
	i := 0
	for k, v := range myTicketFields {
		if strings.HasPrefix(k, "departure") {
			result *= v
			i++
		}
	}

	if i != 6 {
		panic(fmt.Errorf("failed to find 6 fields starting with departure"))
	}

	return result
}
