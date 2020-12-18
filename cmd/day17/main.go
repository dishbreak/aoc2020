package main

import (
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day17.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	cb := func(p *pocketDimension, c lib.Point3D) int {
		state := p.getCubeState(c)
		neighborCount := 0
		for _, neigbhor := range c.Neighbors() {
			neighborCount += p.getCubeState(neigbhor)
		}
		if state == 0 {
			if neighborCount == 3 {
				return 1
			}
			return 0
		}
		if neighborCount == 2 || neighborCount == 3 {
			return 1
		}
		return 0
	}

	p := newPocketDimension(input, cb)

	result := 0
	for i := 0; i < 6; i++ {
		result = p.increment()
	}

	return result
}

func part2(input []string) int {
	cb := func(p *pocketDimension4d, c lib.Point4D) int {
		state := p.getCubeState(c)
		neighborCount := 0
		for _, neigbhor := range c.Neighbors() {
			neighborCount += p.getCubeState(neigbhor)
		}
		if state == 0 {
			if neighborCount == 3 {
				return 1
			}
			return 0
		}
		if neighborCount == 2 || neighborCount == 3 {
			return 1
		}
		return 0
	}

	p := newPocketDimension4d(input, cb)

	result := 0
	for i := 0; i < 6; i++ {
		result = p.increment()
	}
	return result
}
