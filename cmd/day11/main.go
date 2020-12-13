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
}

type Game struct {
	Board      map[image.Point]string
	rows, cols int
}

func NewGame(input []string) *Game {
	board := make(map[image.Point]string)

	for i, row := range input {
		for j, symbol := range row {
			board[image.Point{i, j}] = string(symbol)
		}
	}

	return &Game{
		Board: board,
		rows:  len(input),
		cols:  len(input[0]),
	}
}

func neighbor8(p image.Point) []image.Point {
	return []image.Point{
		image.Point{p.X - 1, p.Y - 1},
		image.Point{p.X - 1, p.Y},
		image.Point{p.X - 1, p.Y + 1},
		image.Point{p.X, p.Y - 1},
		image.Point{p.X, p.Y + 1},
		image.Point{p.X + 1, p.Y - 1},
		image.Point{p.X + 1, p.Y},
		image.Point{p.X + 1, p.Y + 1},
	}
}

func (g *Game) adjacentOccupied(p image.Point) int {
	occupied := 0
	for _, neighbor := range neighbor8(p) {
		if s, ok := g.Board[neighbor]; ok && s == "#" {
			occupied++
		}
	}
	return occupied
}

func (g *Game) Increment() bool {
	changed := false
	board := make(map[image.Point]string)
	for k, v := range g.Board {
		neighbors := g.adjacentOccupied(k)
		if v == "L" && neighbors == 0 {
			board[k] = "#"
			changed = true
		} else if v == "#" && neighbors > 3 {
			board[k] = "L"
			changed = true
		} else {
			board[k] = v
		}
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
	g := NewGame(input)

	for i := 0; g.Increment(); i++ {
		// fmt.Printf("--- Round %d\n", i)
		// fmt.Println(g)
		// fmt.Printf("Occupied: %d\n", g.OccupiedSeats())
	}
	return g.OccupiedSeats()
}
