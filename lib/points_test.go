package lib_test

import (
	"testing"

	"github.com/dishbreak/aoc2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestNeighbors(t *testing.T) {
	p := lib.Point3D{}
	expected := []lib.Point3D{
		lib.Point3D{X: -1, Y: -1, Z: -1},
		lib.Point3D{X: -1, Y: -1, Z: 0},
		lib.Point3D{X: -1, Y: -1, Z: 1},
		lib.Point3D{X: -1, Y: 0, Z: -1},
		lib.Point3D{X: -1, Y: 0, Z: 0},
		lib.Point3D{X: -1, Y: 0, Z: 1},
		lib.Point3D{X: -1, Y: 1, Z: -1},
		lib.Point3D{X: -1, Y: 1, Z: 0},
		lib.Point3D{X: -1, Y: 1, Z: 1},
		lib.Point3D{X: 0, Y: -1, Z: -1},
		lib.Point3D{X: 0, Y: -1, Z: 0},
		lib.Point3D{X: 0, Y: -1, Z: 1},
		lib.Point3D{X: 0, Y: 0, Z: -1},
		lib.Point3D{X: 0, Y: 0, Z: 1},
		lib.Point3D{X: 0, Y: 1, Z: -1},
		lib.Point3D{X: 0, Y: 1, Z: 0},
		lib.Point3D{X: 0, Y: 1, Z: 1},
		lib.Point3D{X: 1, Y: -1, Z: -1},
		lib.Point3D{X: 1, Y: -1, Z: 0},
		lib.Point3D{X: 1, Y: -1, Z: 1},
		lib.Point3D{X: 1, Y: 0, Z: -1},
		lib.Point3D{X: 1, Y: 0, Z: 0},
		lib.Point3D{X: 1, Y: 0, Z: 1},
		lib.Point3D{X: 1, Y: 1, Z: -1},
		lib.Point3D{X: 1, Y: 1, Z: 0},
		lib.Point3D{X: 1, Y: 1, Z: 1}}
	assert.Equal(t, expected, p.Neighbors())
}
