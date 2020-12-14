package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day14.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

var maskMatcher = regexp.MustCompile(`^mask = ([01X]+)$`)
var memMatcher = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

type Mask struct {
	Adds      []int64
	Subtracts []int64
}

func NewMask(maskRepr string) *Mask {
	adds := make([]int64, 0)
	subtracts := make([]int64, 0)

	for idx, symbol := range maskRepr {
		switch symbol {
		case 'X':
			continue
		case '1':
			adds = append(adds, 1<<(35-idx))
		case '0':
			subtracts = append(subtracts, 1<<(35-idx))
		}
	}

	return &Mask{
		Adds:      adds,
		Subtracts: subtracts,
	}
}

func (m *Mask) WriteTo(mem map[int64]int64, addr, value int64) {
	for _, adder := range m.Adds {
		if adder&value == 0 {
			value += adder
		}
	}
	for _, subtractor := range m.Subtracts {
		if subtractor&value != 0 {
			value -= subtractor
		}
	}

	mem[addr] = value
}

func part1(input []string) int64 {
	result := make(map[int64]int64)
	var mask *Mask

	for _, instr := range input {
		maskMatches := maskMatcher.FindAllStringSubmatch(instr, -1)
		if len(maskMatches) != 0 {
			mask = NewMask(maskMatches[0][1])
		}
		memMatches := memMatcher.FindAllStringSubmatch(instr, -1)
		if len(memMatches) != 0 {
			// we are throwing away errors here because regexes already
			// validated what we need.
			addr, _ := strconv.Atoi(memMatches[0][1])
			value, _ := strconv.ParseInt(memMatches[0][2], 10, 64)
			mask.WriteTo(result, int64(addr), int64(value))
		}
	}

	var sum int64
	for _, v := range result {
		sum += v
	}

	return sum
}

func part2(input []string) int {
	return 0
}
