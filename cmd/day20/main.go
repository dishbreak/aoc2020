package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

type edge int

const (
	north edge = 0
	east  edge = 1
	south edge = 2
	west  edge = 3
)

type tile struct {
	id    int
	bits  int
	edges map[edge]int
	raw   []string
}

func main() {
	input, err := lib.GetInputAsSections("inputs/day20.txt")
	if err != nil {
		panic(err)
	}

	tiles := make([]*tile, len(input))
	for idx, section := range input {
		tiles[idx] = toTile(section)
	}

	fmt.Printf("Part 1: %d\n", part1(tiles))
	fmt.Printf("Part 2: %d\n", part2(tiles))
}

func rev(input int, bits int) int {
	rev := 0
	counter := (1 << 10) - 1
	for counter > 0 {
		masked := input & 1
		rev = rev | masked
		rev = rev << 1
		input = input >> 1
		counter = counter >> 1
	}
	rev = rev >> 1
	return rev
}

func toTile(input []string) *tile {
	result := &tile{
		edges: map[edge]int{
			north: 0,
			east:  0,
			south: 0,
			west:  0,
		},
	}

	// parse id out of the first line.
	parts := strings.Split(input[0], " ")
	id, _ := strconv.Atoi(strings.Trim(parts[1], ":"))
	result.id = id
	result.raw = input[1:]

	for lineNo := 1; lineNo < len(input); lineNo++ {
		line := input[lineNo]
		if line[0] == '#' {
			result.edges[west] = result.edges[west] | 1<<(lineNo-1)
		}
		if line[len(line)-1] == '#' {
			result.edges[east] = result.edges[east] | 1<<(lineNo-1)
		}
	}

	northEdge := input[1]
	southEdge := input[len(input)-1]
	for i := 0; i < len(northEdge); i++ {
		if northEdge[i] == '#' {
			result.edges[north] = result.edges[north] | 1<<i
		}
		if southEdge[i] == '#' {
			result.edges[south] = result.edges[south] | 1<<i
		}
	}

	result.bits = len(northEdge)
	return result
}

func (t *tile) getEdges() []int {
	return []int{
		t.edges[north],
		t.edges[east],
		t.edges[south],
		t.edges[west],
		rev(t.edges[north], t.bits),
		rev(t.edges[east], t.bits),
		rev(t.edges[south], t.bits),
		rev(t.edges[west], t.bits),
	}
}

func (t *tile) rotate() *tile {
	o := &tile{
		id:  t.id,
		raw: t.raw,
		edges: map[edge]int{
			north: rev(t.edges[west], t.bits),
			east:  t.edges[north],
			south: rev(t.edges[east], t.bits),
			west:  t.edges[south],
		},
		bits: t.bits,
	}

	return o
}

func (t *tile) flipHorizontal() *tile {
	o := &tile{
		id:  t.id,
		raw: t.raw,
		edges: map[edge]int{
			north: t.edges[south],
			west:  rev(t.edges[west], t.bits),
			east:  rev(t.edges[east], t.bits),
			south: t.edges[north],
		},
		bits: t.bits,
	}
	return o
}

func (t *tile) flipVertical() *tile {
	o := &tile{
		id:  t.id,
		raw: t.raw,
		edges: map[edge]int{
			north: rev(t.edges[north], t.bits),
			east:  t.edges[west],
			south: rev(t.edges[south], t.bits),
			west:  t.edges[east],
		},
		bits: t.bits,
	}
	return o
}

func part1(input []*tile) int {
	// start by remapping all the possible edge values.
	// this will produce a map that contains all tiles with the given edge.
	edgesForTile := make(map[int][]*tile)
	for _, t := range input {
		for _, e := range t.getEdges() {
			l, ok := edgesForTile[e]
			if !ok {
				l = make([]*tile, 0)
			}
			l = append(l, t)
			edgesForTile[e] = l
		}
	}

	// our corner tiles include two edges that cannot match to anything.
	// in the previous map, if a given edge maps to just 1 tile, the tile
	// has an unmatched edge with the given value
	unmatchedEdges := make(map[int][]int)

	for edge, tiles := range edgesForTile {
		if len(tiles) == 1 {
			t, ok := unmatchedEdges[tiles[0].id]
			if !ok {
				t = make([]int, 0)
			}
			t = append(t, edge)
			unmatchedEdges[tiles[0].id] = t
		}
	}

	// now, search the dataset for tiles with 4 unmatched edges.
	// 4 you say?! This is because in order for a tile to have 2
	// unmatched edges, no rotation or translation (which generates a reverse)
	// will make more edges match.
	matches := 0
	acc := 1
	for tileId, ue := range unmatchedEdges {
		if len(ue) == 4 {
			acc *= tileId
			matches++
		}
	}

	if matches != 4 {
		log.Fatalf("incorrect number of matches! expected 4, got %d", matches)
		return -1
	}

	return acc
}

func part2(input []*tile) int {
	return 0
}
