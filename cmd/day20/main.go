package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

func main() {
	input, err := lib.GetInputAsSections("inputs/day20.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

type tileVariant int

const (
	Rotate0 tileVariant = iota
	Rotate90
	Rotate180
	Rotate270
	FlipHorizontal
	FlipVertical
)

type tileQualifier struct {
	id      int
	variant tileVariant
}

type edgeValue int
type edge struct {
	value      edgeValue
	complement edgeValue
}

func (e edge) flip() edge {
	return edge{
		value:      e.complement,
		complement: e.value,
	}
}

type tile struct {
	qualifier tileQualifier
	inUse     *bool
	north     edge
	east      edge
	south     edge
	west      edge
}

type tileSet map[edgeValue][]*tile
type tileBox struct {
	inventory   map[tileQualifier]*tile
	byNorthEdge tileSet
	byEastEdge  tileSet
	byWestEdge  tileSet
	bySouthEdge tileSet
}

func (t *tileBox) registerTile(newTile *tile) {
	t.inventory[newTile.qualifier] = newTile
	if _, ok := t.byNorthEdge[newTile.north.value]; !ok {
		t.byNorthEdge[newTile.north.value] = make([]*tile, 0)
	}
	t.byNorthEdge[newTile.north.value] = append(t.byNorthEdge[newTile.north.value], newTile)

	if _, ok := t.bySouthEdge[newTile.south.value]; !ok {
		t.bySouthEdge[newTile.south.value] = make([]*tile, 0)
	}
	t.bySouthEdge[newTile.south.value] = append(t.bySouthEdge[newTile.south.value], newTile)

	if _, ok := t.byEastEdge[newTile.east.value]; !ok {
		t.byEastEdge[newTile.east.value] = make([]*tile, 0)
	}
	t.byEastEdge[newTile.east.value] = append(t.byEastEdge[newTile.east.value], newTile)

	if _, ok := t.byWestEdge[newTile.west.value]; !ok {
		t.byWestEdge[newTile.west.value] = make([]*tile, 0)
	}
	t.byWestEdge[newTile.west.value] = append(t.byWestEdge[newTile.west.value], newTile)
}

func reverseNumber(input edgeValue) edgeValue {
	var result edgeValue
	for input > 0 {
		result <<= 1
		if input&1 == 1 {
			result ^= 1
		}
		input >>= 1
	}
	return result
}

func (t *tileBox) registerInput(input []string) {
	inUse := false
	titleParts := strings.Fields(input[0])
	tileString := strings.TrimSuffix(titleParts[1], ":")

	tileNumber, err := strconv.Atoi(tileString)
	if err != nil {
		panic(err)
	}

	var north, east, south, west edge
	tileContents := input[1:]
	for i, row := range tileContents {
		for j, col := range row {
			switch col {
			case '#':
				if i == 0 {
					north.value = north.value | 1<<j
				}
				if i == len(tileContents)-1 {
					south.value = south.value | 1<<(len(row)-1-j)
				}
				if j == 0 {
					west.value = west.value | 1<<(len(tileContents)-1-i)
				}
				if j == len(row)-1 {
					east.value = east.value | 1<<i
				}
			}
		}
	}

	north.complement = reverseNumber(north.value)
	south.complement = reverseNumber(south.value)
	east.complement = reverseNumber(east.value)
	west.complement = reverseNumber(west.value)

	baseTile := &tile{
		qualifier: tileQualifier{
			id:      tileNumber,
			variant: Rotate0,
		},
		inUse: &inUse,
		north: north,
		east:  east,
		south: south,
		west:  west,
	}

	t.registerTile(baseTile)

	// keep rotating the tile by 90 degrees and storing the result
	previousTile := baseTile
	for i := Rotate90; i < FlipHorizontal; i++ {
		nextTile := &tile{
			qualifier: tileQualifier{
				id:      tileNumber,
				variant: i,
			},
			inUse: &inUse,
			north: previousTile.west,
			east:  previousTile.north,
			south: previousTile.east,
			west:  previousTile.south,
		}
		t.registerTile(nextTile)
		previousTile = nextTile
	}

	t.registerTile(&tile{
		qualifier: tileQualifier{
			id:      tileNumber,
			variant: FlipHorizontal,
		},
		inUse: &inUse,
		north: baseTile.south,
		south: baseTile.north,
		west:  baseTile.west.flip(),
		east:  baseTile.east.flip(),
	})

	t.registerTile(&tile{
		qualifier: tileQualifier{
			id:      tileNumber,
			variant: FlipVertical,
		},
		inUse: &inUse,
		west:  baseTile.east,
		east:  baseTile.west,
		north: baseTile.north.flip(),
		south: baseTile.south.flip(),
	})

}

func newTileBox(input [][]string) *tileBox {
	t := &tileBox{
		inventory:   make(map[tileQualifier]*tile, len(input)*5),
		byNorthEdge: make(tileSet, len(input)*5),
		byEastEdge:  make(tileSet, len(input)*5),
		byWestEdge:  make(tileSet, len(input)*5),
		bySouthEdge: make(tileSet, len(input)*5),
	}

	for _, blob := range input {
		t.registerInput(blob)
	}
	return t
}

func part1(input [][]string) int {
	return 0
}

func part2(input [][]string) int {
	return 0
}
