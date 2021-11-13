package main

import (
	"fmt"
)

func main() {
	input := "583976241"

	fmt.Printf("Part 1: %s\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) string {
	g := newShellGame(input)
	for i := 0; i < 100; i++ {
		g.playRound()
	}

	return g.String()
}

func part2(input string) int {
	g := newCrabShellGame(input)

	for i := 0; i < 10000000; i++ {
		g.playRound()
	}

	c := g.getCup(1)
	return c.Next.Data * c.Next.Next.Data
}
