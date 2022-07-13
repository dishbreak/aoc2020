package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []*tile{}

var testFile = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`

func getTestData() []*tile {
	d := make([]*tile, 0)
	for _, blob := range strings.Split(testFile, "\n\n") {
		d = append(d, toTile(strings.Split(blob, "\n")))
	}
	return d
}

func TestPart1(t *testing.T) {
	d := getTestData()
	assert.Equal(t, 20899048083289, part1(d))
}

func TestPart2(t *testing.T) {
	d := getTestData()
	assert.Equal(t, 273, part2(d))
}

func TestToTile(t *testing.T) {
	input := []string{
		"Tile 2311:",
		"..##.#..#.",
		"##..#.....",
		"#...##..#.",
		"####.#...#",
		"##.##.###.",
		"##...#.###",
		".#.#.#..##",
		"..#....#..",
		"###...#.#.",
		"..###..###",
	}

	myTile := toTile(input)
	expected := &tile{
		id: 2311,
		edges: map[edge]int{
			north: 300,
			east:  616,
			south: 924,
			west:  318,
		},
		raw: []string{
			"..##.#..#.",
			"##..#.....",
			"#...##..#.",
			"####.#...#",
			"##.##.###.",
			"##...#.###",
			".#.#.#..##",
			"..#....#..",
			"###...#.#.",
			"..###..###",
		},
		bits: 10,
	}

	assert.Equal(t, expected, myTile)
}

func TestRev(t *testing.T) {
	type testData struct {
		expected int
		input    int
	}

	tests := map[string]testData{
		"basic": {
			0b0010010011,
			0b1100100100,
		},
		"bug in rotate": {
			498,
			318,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := rev(test.input, 10)
			assert.Equal(t, test.expected, actual, "expected %0bd but got %0bd", test.expected, actual)
		})
	}

}
