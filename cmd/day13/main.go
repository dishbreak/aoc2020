package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day13.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type BusCalculator struct {
	startTime int
	busIds    []int
	maxId     int
}

func NewBusCalculator(input []string) (*BusCalculator, error) {
	startTime, err := strconv.Atoi(input[0])
	if err != nil {
		return nil, err
	}

	maxId := 0
	busIds := make([]int, 0)
	for _, id := range strings.Split(input[1], ",") {
		if id == "x" {
			continue
		}
		if parsed, err := strconv.Atoi(id); err == nil {
			if maxId < parsed {
				maxId = parsed
			}
			busIds = append(busIds, parsed)
		} else {
			return nil, err
		}
	}

	return &BusCalculator{
		startTime: startTime,
		busIds:    busIds,
		maxId:     maxId,
	}, nil
}

func (b *BusCalculator) FindNextBus() (int, int) {
	for i := b.startTime; i <= b.startTime+b.maxId; i++ {
		for _, busId := range b.busIds {
			if i%busId == 0 {
				return busId, i - b.startTime
			}
		}
	}
	return -1, -1
}

func part1(input []string) int {
	calc, err := NewBusCalculator(input)
	if err != nil {
		panic(err)
	}

	id, timeToBus := calc.FindNextBus()
	return id * timeToBus
}

func part2(input []string) int {
	return 0
}
