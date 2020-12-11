package main

import (
	"bytes"
	"fmt"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day11.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

type Space struct {
	Floor bool
	State int
}

// adapted from https://golang.org/doc/play/life.go
type Field struct {
	s          [][]bool
	rows, cols int
	f          [][]bool
}

func NewField(rows, cols int, input []string) *Field {
	s := make([][]bool, rows)
	for i := range s {
		s[i] = make([]bool, cols)
	}

	f := make([][]bool, rows)
	for i := range f {
		f[i] = make([]bool, cols)
	}

	for i, row := range input {
		for j, col := range row {
			f[i][j] = col == '.'
		}
	}

	return &Field{s: s, f: f, rows: rows, cols: cols}
}

func (f *Field) IsFilled(row, col int) bool {
	if row < 0 || col < 0 || row >= f.rows || col >= f.cols {
		return false
	}
	return f.s[row][col]
}

func (f *Field) FilledNeighbors(row, col int) int {
	filledNeigbhors := 0
	neighborCoords := [8][2]int{
		[2]int{row - 1, col - 1},
		[2]int{row - 1, col},
		[2]int{row - 1, col + 1},
		[2]int{row, col - 1},
		[2]int{row, col + 1},
		[2]int{row + 1, col - 1},
		[2]int{row + 1, col},
		[2]int{row + 1, col + 1},
	}

	for _, point := range neighborCoords {
		if f.IsFilled(point[0], point[1]) {
			filledNeigbhors++
		}
	}

	return filledNeigbhors
}

type Game struct {
	a          *Field
	b          *Field
	rows, cols int
}

func NewGame(input []string) *Game {
	rows := len(input)
	cols := len(input[0])
	return &Game{
		rows: rows,
		cols: cols,
		a:    NewField(rows, cols, input),
		b:    NewField(rows, cols, input),
	}
}

func (g *Game) Increment() bool {
	changed := false
	for i, row := range g.a.s {
		for j := range row {
			if g.a.f[i][j] {
				continue
			}
			neighbors := g.a.FilledNeighbors(i, j)
			if g.a.s[i][j] && neighbors >= 4 {
				g.b.s[i][j] = false
				changed = true
			} else if !g.a.s[i][j] && neighbors == 0 {
				g.b.s[i][j] = true
				changed = true
			} else {
				g.b.s[i][j] = g.a.s[i][j]
			}
		}
	}

	g.a, g.b = g.b, g.a
	return changed
}

func (g *Game) OccupiedSeats() int {
	occupied := 0
	for i, row := range g.a.s {
		for j := range row {
			if g.a.s[i][j] {
				occupied++
			}
		}
	}
	return occupied
}
func (g *Game) String() string {
	var b bytes.Buffer
	for i, row := range g.a.s {
		for j := range row {
			if g.a.f[i][j] {
				b.WriteByte('.')
			} else if g.a.s[i][j] {
				b.WriteByte('#')
			} else {
				b.WriteByte('L')
			}
		}
		b.WriteByte('\n')
	}

	return b.String()
}

func part1(input []string) int {
	// build out a grid surrounded by floor. This makes boundary conditions easier to deal with.
	game := NewGame(input)

	for i := 0; game.Increment(); i++ {
		// fmt.Printf("+++ Round %d +++\n", i)
		// fmt.Println(game)
	}

	return game.OccupiedSeats()
}
