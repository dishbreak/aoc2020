package main

import (
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

type direction int

const (
	w direction = iota
	nw
	ne
	e
	se
	sw
)

func tokenize(input string) []direction {
	result := make([]direction, 0)

	buf := strings.NewReader(input)

	for r, _, err := buf.ReadRune(); err == nil; r, _, err = buf.ReadRune() {
		var d direction
		switch r {
		case 'e':
			d = e
		case 'w':
			d = w
		case 'n':
			s, _, _ := buf.ReadRune()
			switch s {
			case 'e':
				d = ne
			case 'w':
				d = nw
			}
		case 's':
			s, _, _ := buf.ReadRune()
			switch s {
			case 'e':
				d = se
			case 'w':
				d = sw
			}
		}
		result = append(result, d)
	}
	return result
}

var directionToVector map[direction]lib.Point3D = map[direction]lib.Point3D{
	w:  {X: -1, Y: 0, Z: 1},
	nw: {X: 0, Y: -1, Z: 1},
	ne: {X: 1, Y: -1, Z: 0},
	e:  {X: 1, Y: 0, Z: -1},
	se: {X: 0, Y: 1, Z: -1},
	sw: {X: -1, Y: 1, Z: 0},
}

func traverse(directions []direction) lib.Point3D {
	result := lib.Point3D{}
	for _, d := range directions {
		result = result.Add(directionToVector[d])
	}
	return result
}

type lobbyFloor struct {
	tiles map[lib.Point3D]int
}

func (l *lobbyFloor) flipTile(instructions string) {
	d := tokenize(instructions)
	p := traverse(d)
	if _, ok := l.tiles[p]; !ok {
		l.tiles[p] = 0
	}
	l.tiles[p]++
}

func (l *lobbyFloor) normalize() {
	for p, val := range l.tiles {
		val = val % 2
		if val == 0 {
			l.tiles[p] = 0
		} else {
			l.tiles[p] = 1
		}

		for _, v := range directionToVector {
			d := p.Add(v)
			if _, ok := l.tiles[d]; !ok {
				l.tiles[d] = 0
			}
		}
	}
}

func (l *lobbyFloor) countBlackTiles() int {
	result := 0
	for _, val := range l.tiles {
		result += val % 2
	}
	return result
}

func (l *lobbyFloor) countNeighbors(p lib.Point3D) int {
	acc := 0
	for _, v := range directionToVector {
		acc += l.tiles[p.Add(v)]
	}
	return acc
}

func (l *lobbyFloor) advanceDay() {
	l.normalize()

	ntiles := make(map[lib.Point3D]int)

	for p, val := range l.tiles {
		b := l.countNeighbors(p)
		switch val {
		case 0:
			if b == 2 {
				ntiles[p] = 1
			} else {
				ntiles[p] = 0
			}
		case 1:
			if b == 0 || b > 2 {
				ntiles[p] = 0
			} else {
				ntiles[p] = 1
			}
		}
	}

	l.tiles = ntiles
}
