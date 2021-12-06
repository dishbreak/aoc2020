package lib

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	input := []string{
		"##..#.",
		"#..###",
		"..#..#",
		"...###",
		"#...##",
		"###...",
	}

	output := strings.Join([]string{
		"##..##",
		"#....#",
		"#..#..",
		"..#.#.",
		".##.##",
		".####.",
		"",
	}, "\n")

	m := NewMatrix(input)
	m.Rotate()
	assert.Equal(t, output, m.String())
}

func TestNewMatrixWithoutFrame(t *testing.T) {
	start := []string{
		"##..#.",
		"#..###",
		"..#..#",
		"...###",
		"#...##",
		"###...",
	}

	result := [][]byte{
		[]byte("..##"),
		[]byte(".#.."),
		[]byte("..##"),
		[]byte("...#"),
	}

	m := NewMatrixWithoutFrame(start)
	assert.Equal(t, result, m.d)
}
