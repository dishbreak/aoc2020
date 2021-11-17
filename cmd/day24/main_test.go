package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"sesenwnenenewseeswwswswwnenewsewsw",
	"neeenesenwnwwswnenewnwwsewnenwseswesw",
	"seswneswswsenwwnwse",
	"nwnwneseeswswnenewneswwnewseswneseene",
	"swweswneswnenwsewnwneneseenw",
	"eesenwseswswnenwswnwnwsewwnwsene",
	"sewnenenenesenwsewnenwwwse",
	"wenwwweseeeweswwwnwwe",
	"wsweesenenewnwwnwsenewsenwwsesesenwne",
	"neeswseenwwswnwswswnw",
	"nenwswwsewswnenenewsenwsenwnesesenew",
	"enewnwewneswsewnwswenweswnenwsenwsw",
	"sweneswneswneneenwnewenewwneswswnese",
	"swwesenesewenwneswnwwneseswwne",
	"enesenwswwswneneswsenwnewswseenwsese",
	"wnwnesenesenenwwnenwsewesewsesesew",
	"nenewswnwewswnenesenwnesewesw",
	"eneswnwswnwsenenwnwnwwseeswneewsenese",
	"neswnwewnwnwseenwseesewsenwsweewe",
	"wseweeenwnesenwwwswnew",
}

func TestPart1(t *testing.T) {
	assert.Equal(t, 10, part1(input))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 2208, part2(input))
}
