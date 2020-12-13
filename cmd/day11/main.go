package main

import (
	"bytes"
	"fmt"
	"image"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day11.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type Game struct {
	Board      map[image.Point]string
	rows, cols int
	callback   func(*Game, image.Point, string) (string, bool)
}

func NewGame(input []string, callback func(*Game, image.Point, string) (string, bool)) *Game {
	board := make(map[image.Point]string)

	for i, row := range input {
		for j, symbol := range row {
			board[image.Point{i, j}] = string(symbol)
		}
	}

	return &Game{
		Board:    board,
		rows:     len(input),
		cols:     len(input[0]),
		callback: callback,
	}
}

var neighbor8 = []image.Point{
	image.Point{-1, -1},
	image.Point{-1, 0},
	image.Point{-1, 1},
	image.Point{0, -1},
	image.Point{0, 1},
	image.Point{1, -1},
	image.Point{1, 0},
	image.Point{1, 1},
}

func (g *Game) adjacentOccupied(p image.Point) int {
	occupied := 0
	for _, neighbor := range neighbor8 {
		if s, ok := g.Board[p.Add(neighbor)]; ok && s == "#" {
			occupied++
		}
	}
	return occupied
}

func (g *Game) seenOccupied(p image.Point) int {
	seen := 0
	for _, neighbor := range neighbor8 {
		j := p.Add(neighbor)
		for s, ok := g.Board[j]; ok; s, ok = g.Board[j] {
			if s == "#" {
				seen++
				break
			} else if s == "L" {
				break
			}
			j = j.Add(neighbor)
		}
	}
	return seen
}

func (g *Game) Increment() bool {
	changed := false
	board := make(map[image.Point]string)
	for k, v := range g.Board {
		result, updated := g.callback(g, k, v)
		changed = changed || updated
		board[k] = result
	}

	g.Board = board
	return changed
}

func (g *Game) OccupiedSeats() int {
	occupied := 0
	for _, v := range g.Board {
		if v == "#" {
			occupied++
		}
	}
	return occupied
}

func (g *Game) String() string {
	var b bytes.Buffer
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			b.WriteString(g.Board[image.Point{i, j}])
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func part1(input []string) int {
	f := func(g *Game, p image.Point, s string) (string, bool) {
		neighbors := g.adjacentOccupied(p)
		if s == "L" && neighbors == 0 {
			return "#", true
		} else if s == "#" && neighbors > 3 {
			return "L", true
		} else {
			return s, false
		}
	}

	g := NewGame(input, f)

	for g.Increment() {
	}

	return g.OccupiedSeats()
}

func part2(input []string) int {
	f := func(g *Game, p image.Point, s string) (string, bool) {
		neighbors := g.seenOccupied(p)
		if s == "L" && neighbors == 0 {
			return "#", true
		} else if s == "#" && neighbors > 4 {
			return "L", true
		} else {
			return s, false
		}
	}

	g := NewGame(input, f)
	for g.Increment() {
	}

	return g.OccupiedSeats()
}
