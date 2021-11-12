package main

import (
	"strconv"
	"strings"
)

type shellGame struct {
	cups       []int
	currentCup int
	max        int
	min        int
	startCup   int
}

func newShellGame(input string) *shellGame {
	g := &shellGame{
		cups: make([]int, len(input)),
		min:  10,
		max:  -1,
	}

	for idx, char := range strings.Split(input, "") {
		parsed, _ := strconv.Atoi(char)
		if parsed < g.min {
			g.min = parsed
		}
		if parsed > g.max {
			g.max = parsed
		}
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
	if idx == len(g.cups) {
		g.cups = append(g.cups, val)
		return
	}

	idx = g.normalize(idx)

	g.cups = append(g.cups[:idx+1], g.cups[idx:]...)
	g.cups[idx] = val
}

func (g *shellGame) decrement(v int) int {
	v--
	if v < g.min {
		v = g.max
	}
	return v
}

func (g *shellGame) playRound() {
	target := g.decrement(g.cups[0])

	// make a bit vector to keep track of which cups are removed
	bv := make([]int, 10)

	// make a buffer to hold onto the cups while they're out of play
	buf := make([]int, 3)
	copy(buf, g.cups[1:4])
	g.cups = append(g.cups[0:1], g.cups[4:]...)

	for _, val := range buf {
		bv[val] = 1
	}

	// keep decrementing the target until the target is no longer one of the removed cups.
	for bv[target] == 1 {
		target = g.decrement(target)
	}

	// find the index with the desired cup value
	i := 1
	for ; i < len(g.cups); i++ {
		if g.cups[i] == target {
			break
		}
	}

	// if the target cup is at the end of the slice, this is easy!
	if i == len(g.cups)-1 {
		g.cups = append(g.cups, buf...)
	} else { // if not, well...
		// lop off the slice following the target cup, store it in scratch.
		scratch := make([]int, len(g.cups)-i-1)
		copy(scratch, g.cups[i+1:])
		g.cups = g.cups[:i+1]

		// add the removed cups back to the game.
		g.cups = append(g.cups[:i+1], buf...)
		// add the lopped off cups afterwards.
		g.cups = append(g.cups, scratch...)
	}

	// rotate the slice so that the next current cup is at the head of the list.
	g.cups = append(g.cups[1:], g.cups[0])
}

func (g *shellGame) String() string {
	idx := 0
	for ; g.cups[idx] != 1; idx++ {
		continue
	}

	buf := strings.Builder{}
	for iter := g.normalize(idx + 1); iter != idx; iter = g.normalize(iter + 1) {
		buf.Write([]byte(strconv.Itoa(g.cups[iter])))
	}
	return buf.String()
}
