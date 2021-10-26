package main

import (
	"strconv"
	"strings"
)

type shellGame struct {
	cups []int
}

func newShellGame(input string) *shellGame {
	g := &shellGame{
		cups: make([]int, len(input)),
	}

	for idx, char := range strings.Split(input, "") {
		parsed, _ := strconv.Atoi(char)
		g.cups[idx] = parsed
	}

	return g
}

func (g *shellGame) normalize(idx int) int {
	idx = idx % len(g.cups)
	if idx < 0 {
		idx = len(g.cups) + idx
	}
	return idx
}

func (g *shellGame) removeAt(idx int) (val int) {

	idx = g.normalize(idx)
	val = g.cups[idx]

	if idx == 0 {
		g.cups = g.cups[1:]
		return
	}

	if idx == len(g.cups)-1 {
		g.cups = g.cups[:idx]
		return
	}

	g.cups = append(g.cups[:idx], g.cups[idx+1:]...)
	return
}

func (g *shellGame) insertAt(val, idx int) {
	idx = g.normalize(idx)

	g.cups = append(g.cups[:idx+1], g.cups[idx:]...)
	g.cups[idx] = val
}
