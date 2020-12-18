package main

import "github.com/dishbreak/aoc2020/lib"

type pocketDimension4d struct {
	cube          map[lib.Point4D]int
	min           lib.Point4D
	max           lib.Point4D
	rulesCallback func(*pocketDimension4d, lib.Point4D) int
}

func newPocketDimension4d(input []string, callback func(*pocketDimension4d, lib.Point4D) int) *pocketDimension4d {
	// presuming a rectangular input
	bounds := lib.Point4D{
		X: len(input),
		Y: len(input[0]),
		Z: 0,
	}

	p := &pocketDimension4d{
		cube:          make(map[lib.Point4D]int),
		min:           lib.Point4D{X: -1, Y: -1, Z: -1, W: -1},
		max:           bounds.Add(lib.Point4D{X: 1, Y: 1, Z: 1, W: 1}),
		rulesCallback: callback,
	}

	for i, line := range input {
		for j, symbol := range line {
			c := lib.Point4D{X: i, Y: j, Z: 0, W: 0}
			if symbol == '#' {
				p.cube[c] = 1
				continue
			}
			p.cube[c] = 0
		}
	}

	return p
}

func (p *pocketDimension4d) increment() int {
	activeCubes := 0
	ncube := make(map[lib.Point4D]int)

	for i := p.min.X; i <= p.max.X; i++ {
		for j := p.min.Y; j <= p.max.Y; j++ {
			for k := p.min.Z; k <= p.max.Z; k++ {
				for l := p.min.W; l <= p.max.W; l++ {
					c := lib.Point4D{X: i, Y: j, Z: k, W: l}
					state := p.rulesCallback(p, c)
					activeCubes += state
					ncube[c] = state
				}
			}
		}
	}

	p.cube = ncube
	p.min = p.min.Sub(lib.Point4D{X: 1, Y: 1, Z: 1, W: 1})
	p.max = p.max.Add(lib.Point4D{X: 1, Y: 1, Z: 1, W: 1})

	return activeCubes
}

func (p *pocketDimension4d) getCubeState(c lib.Point4D) int {
	if state, ok := p.cube[c]; ok {
		return state
	}
	return 0
}
