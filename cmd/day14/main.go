package main

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
	"github.com/go-errors/errors"
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

type MaskV1 struct {
	Adds      []int64
	Subtracts []int64
}

func NewMaskV1(maskRepr string) Mask {
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

	return &MaskV1{
		Adds:      adds,
		Subtracts: subtracts,
	}
}

func (m *MaskV1) WriteTo(mem map[int64]int64, addr, value int64) {
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

type MaskV2 struct {
	Floaters  []int64
	Adds      []int64
	Subtracts []int64
}

func NewMaskV2(maskRepr string) Mask {
	floaters := make([]int64, 0)
	adds := make([]int64, 0)
	subtracts := make([]int64, 0)
	for idx, symbol := range maskRepr {
		value := int64(1 << (35 - idx))
		switch symbol {
		case 'X':
			floaters = append(floaters, value)
		case '1':
			adds = append(adds, value)
		case '0':
			subtracts = append(subtracts, value)
		}
	}

	return &MaskV2{floaters, adds, subtracts}
}

type queueEntry struct {
	value int64
	level int
}

func (m *MaskV2) WriteTo(mem map[int64]int64, addr, value int64) {
	for _, adder := range m.Adds {
		if adder&addr == 0 {
			addr += adder
		}
	}
	for _, floater := range m.Floaters {
		if floater&addr != 0 {
			addr -= floater
		}
	}

	l := list.New()
	l.PushBack(queueEntry{addr, 0})

	for l.Len() > 0 {
		e := l.Front()
		v, ok := e.Value.(queueEntry)
		if !ok {
			panic(errors.New("unexpected entry in queue"))
		}
		if v.level == len(m.Floaters) {
			mem[v.value] = value
		} else {
			l.PushBack(queueEntry{v.value, v.level + 1})
			l.PushBack(queueEntry{v.value + m.Floaters[v.level], v.level + 1})
		}
		l.Remove(e)
	}
}

type Mask interface {
	WriteTo(mem map[int64]int64, addr, value int64)
}
type Emulator struct {
	mem         map[int64]int64
	maskFactory func(string) Mask
}

func (e *Emulator) RunInstructions(instructions []string) {
	var mask Mask
	for _, instr := range instructions {
		maskMatches := maskMatcher.FindAllStringSubmatch(instr, -1)
		if len(maskMatches) != 0 {
			mask = e.maskFactory(maskMatches[0][1])
		}
		memMatches := memMatcher.FindAllStringSubmatch(instr, -1)
		if len(memMatches) != 0 {
			// we are throwing away errors here because regexes already
			// validated what we need.
			addr, _ := strconv.Atoi(memMatches[0][1])
			value, _ := strconv.ParseInt(memMatches[0][2], 10, 64)
			mask.WriteTo(e.mem, int64(addr), int64(value))
		}
	}
}

func (e *Emulator) SumValues() int64 {
	var sum int64
	for _, v := range e.mem {
		sum += v
	}

	return sum
}

func part1(input []string) int64 {
	e := &Emulator{
		mem:         make(map[int64]int64),
		maskFactory: NewMaskV1,
	}

	e.RunInstructions(input)
	return e.SumValues()
}

func part2(input []string) int64 {
	e := &Emulator{
		mem:         make(map[int64]int64),
		maskFactory: NewMaskV2,
	}

	e.RunInstructions(input)
	return e.SumValues()
}
