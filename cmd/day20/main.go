package main

import (
	"errors"
	"fmt"
	"image"
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
	img   *lib.Matrix
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

	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	// parse id out of the first line.
	parts := strings.Split(input[0], " ")
	id, _ := strconv.Atoi(strings.Trim(parts[1], ":"))
	result.id = id
	result.raw = input[1:]
	result.img = lib.NewMatrix(result.raw)

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

func (t *tile) Rotate() {
	oldEdges := make(map[edge]int)
	for k, v := range t.edges {
		oldEdges[k] = v
	}

	t.edges[east] = oldEdges[north]
	t.edges[south] = rev(oldEdges[east], t.bits)
	t.edges[west] = oldEdges[south]
	t.edges[north] = rev(oldEdges[west], t.bits)

	t.img.Rotate()
}

func (t *tile) FlipVertical() {
	oldEdges := make(map[edge]int)
	for k, v := range t.edges {
		oldEdges[k] = v
	}

	t.edges[east] = oldEdges[west]
	t.edges[west] = oldEdges[east]
	t.edges[north] = rev(oldEdges[north], t.bits)
	t.edges[south] = rev(oldEdges[south], t.bits)

	t.img.FlipVertical()
}

// map all the possible edge values to the corresponding tiles
// this will produce a map that contains all tiles with the given edge.
func mapEdgesToTile(input []*tile) map[int][]*tile {
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
	return edgesForTile
}

func findCornerTiles(edgesForTile map[int][]*tile) []*tile {
	// our corner tiles include two edges that cannot match to anything.
	// in the previous map, if a given edge maps to just 1 tile, the tile
	// has an unmatched edge with the given value
	unmatchedEdges := make(map[*tile][]int)

	for edge, tiles := range edgesForTile {
		if len(tiles) == 1 {
			t, ok := unmatchedEdges[tiles[0]]
			if !ok {
				t = make([]int, 0)
			}
			t = append(t, edge)
			unmatchedEdges[tiles[0]] = t
		}
	}

	// now, search the dataset for tiles with 4 unmatched edges.
	// 4 you say?! This is because in order for a tile to have 2
	// unmatched edges, no rotation or translation (which generates a reverse)
	// will make more edges match.
	matches := make([]*tile, 0)
	for tilePtr, ue := range unmatchedEdges {
		if len(ue) == 4 {
			matches = append(matches, tilePtr)
		}
	}

	return matches
}

func part1(input []*tile) int {

	edgesForTile := mapEdgesToTile(input)

	matches := findCornerTiles(edgesForTile)

	if len(matches) != 4 {
		log.Fatalf("incorrect number of matches! expected 4, got %d", len(matches))
		return -1
	}

	acc := 1
	for _, tilePtr := range matches {
		acc *= tilePtr.id
	}
	return acc
}

type Rotatable interface {
	Rotate()
	FlipVertical()
}

func flipTilTest(r Rotatable, test func(Rotatable) bool) {
	for i := 0; i < 4 && !test(r); i++ {
		r.Rotate()
	}

	for i := 0; i < 4 && !test(r); i++ {
		if i != 0 {
			r.FlipVertical()
		}
		r.Rotate()
		r.FlipVertical()
	}

	if !test(r) {
		panic(errors.New("failed to find passing rotation"))
	}
}

func validateTileSet(img map[image.Point]*tile) bool {
	type neighborCheck struct {
		v    image.Point
		oppE edge
	}

	neighbors := map[edge]neighborCheck{
		north: {
			v:    image.Pt(0, -1),
			oppE: south,
		},
		south: {
			v:    image.Pt(0, 1),
			oppE: north,
		},
		east: {
			v:    image.Pt(1, 0),
			oppE: west,
		},
		west: {
			v:    image.Pt(-1, 0),
			oppE: east,
		},
	}

	for pt, tile := range img {
		for e, n := range neighbors {
			nPt := pt.Add(n.v)
			nTile, ok := img[nPt]
			if !ok {
				continue
			}
			if nTile.edges[n.oppE] != tile.edges[e] {
				return false
			}
		}
	}
	return true
}

func part2(input []*tile) int {
	edgesForTile := mapEdgesToTile(input)

	matches := findCornerTiles(edgesForTile)

	isNorthwestCorner := func(r Rotatable) bool {
		t, _ := r.(*tile)
		for _, e := range []edge{north, west} {
			neighbors := edgesForTile[t.edges[e]]
			if len(neighbors) > 1 {
				return false
			}
		}
		return true
	}

	start := matches[0]

	flipTilTest(start, isNorthwestCorner)
	tileSet := make(map[image.Point]*tile)
	q := make([]image.Point, 1)
	q[0] = image.Point{0, 0}
	tileSet[q[0]] = start

	findNeighbor := func(t *tile, pt image.Point, e edge) (image.Point, bool) {
		var v image.Point
		var oppE edge
		switch e {
		case south:
			v = image.Pt(0, 1)
			oppE = north
		case east:
			v = image.Pt(1, 0)
			oppE = west
		}

		nPt := pt.Add(v)

		if _, ok := tileSet[nPt]; ok {
			return nPt, false
		}

		neighbors := edgesForTile[t.edges[e]]
		if len(neighbors) == 1 {
			return nPt, false
		}

		var hit *tile
		for _, neighbor := range neighbors {
			if t.id != neighbor.id {
				hit = neighbor
				break
			}
		}

		flipTilTest(hit, func(r Rotatable) bool {
			oT, _ := r.(*tile)
			return t.edges[e] == oT.edges[oppE]
		})

		tileSet[nPt] = hit
		return nPt, true
	}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		n, ok := tileSet[p]
		if !ok {
			continue
		}

		for _, e := range []edge{south, east} {
			if pt, ok := findNeighbor(n, p, e); ok {
				q = append(q, pt)
			}
		}
	}

	if !validateTileSet(tileSet) {
		panic(errors.New("image is invalid"))
	}

	return 0
}
