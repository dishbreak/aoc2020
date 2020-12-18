package main

import "github.com/dishbreak/aoc2020/lib"

type pocketDimension struct {
	cube          map[lib.Point3D]int
	min           lib.Point3D
	max           lib.Point3D
	rulesCallback func(*pocketDimension, lib.Point3D) int
}

func newPocketDimension(input []string, callback func(*pocketDimension, lib.Point3D) int) *pocketDimension {
	// presuming a rectangular input
	bounds := lib.Point3D{
		X: len(input),
		Y: len(input[0]),
		Z: 0,
	}

	p := &pocketDimension{
		cube:          make(map[lib.Point3D]int),
		min:           lib.Point3D{X: -1, Y: -1, Z: -1},
		max:           bounds.Add(lib.Point3D{X: 1, Y: 1, Z: 1}),
		rulesCallback: callback,
	}

	for i, line := range input {
		for j, symbol := range line {
			c := lib.Point3D{X: i, Y: j, Z: 0}
			if symbol == '#' {
				p.cube[c] = 1
				continue
			}
			p.cube[c] = 0
		}
	}

	return p
}

func (p *pocketDimension) increment() int {
	activeCubes := 0
	ncube := make(map[lib.Point3D]int)

	for i := p.min.X; i <= p.max.X; i++ {
		for j := p.min.Y; j <= p.max.Y; j++ {
			for k := p.min.Z; k <= p.max.Z; k++ {
				c := lib.Point3D{X: i, Y: j, Z: k}
				state := p.rulesCallback(p, c)
				activeCubes += state
				ncube[c] = state
			}
		}
	}

	p.cube = ncube
	p.min = p.min.Sub(lib.Point3D{X: 1, Y: 1, Z: 1})
	p.max = p.max.Add(lib.Point3D{X: 1, Y: 1, Z: 1})

	return activeCubes
}

func (p *pocketDimension) getCubeState(c lib.Point3D) int {
	if state, ok := p.cube[c]; ok {
		return state
	}
	return 0
}
