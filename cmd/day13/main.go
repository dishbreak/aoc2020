package main

import (
	"fmt"
	"sort"
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

// BusCalculator reads a schedule and lets us know which bus to take to the
// airport!
type BusCalculator struct {
	startTime int
	busIds    []int
	maxID     int
}

// NewBusCalculator parses a schedule and returns a BusCalculator.
func NewBusCalculator(input []string) (*BusCalculator, error) {
	startTime, err := strconv.Atoi(input[0])
	if err != nil {
		return nil, err
	}

	maxID := 0
	busIds := make([]int, 0)
	for _, id := range strings.Split(input[1], ",") {
		if id == "x" {
			continue
		}
		if parsed, err := strconv.Atoi(id); err == nil {
			if maxID < parsed {
				maxID = parsed
			}
			busIds = append(busIds, parsed)
		} else {
			return nil, err
		}
	}

	return &BusCalculator{
		startTime: startTime,
		busIds:    busIds,
		maxID:     maxID,
	}, nil
}

// FindNextBus determines the next departing bus and the relative departuree
// time.
func (b *BusCalculator) FindNextBus() (int, int) {
	for i := b.startTime; i <= b.startTime+b.maxID; i++ {
		for _, busID := range b.busIds {
			// loop through the busids, and exit when we find that the current
			// timestamp is a multiple of a bus ID.
			if i%busID == 0 {
				return busID, i - b.startTime
			}
		}
	}
	return -1, -1
}

// BusScheduleEntry captures the ID and offset of a departing bus.
type BusScheduleEntry struct {
	ID     int
	Offset int
}

func part1(input []string) int {
	calc, err := NewBusCalculator(input)
	if err != nil {
		panic(err)
	}

	id, timeToBus := calc.FindNextBus()
	return id * timeToBus
}

func (e BusScheduleEntry) TimeToNextDeparture(t int) int {
	mod := t % e.ID
	if mod == 0 {
		return 0
	}
	return e.ID - (t % e.ID)
}

func LoadSchedule(input string) []BusScheduleEntry {
	result := make([]BusScheduleEntry, 0)
	for idx, id := range strings.Split(input, ",") {
		if id == "x" {
			continue
		}
		if parsed, err := strconv.Atoi(id); err == nil {
			result = append(result, BusScheduleEntry{parsed, idx % parsed})
		} else {
			panic(err)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID > result[j].ID
	})
	return result
}

func part2(input []string) int {
	schedule := LoadSchedule(input[1])
	t := 2*schedule[0].ID - schedule[0].Offset
	increment := schedule[0].ID

	for _, entry := range schedule[1:] {
		for entry.TimeToNextDeparture(t) != entry.Offset {
			t += increment
		}
		increment *= entry.ID
	}
	return t
}
