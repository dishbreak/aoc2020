package lib_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dishbreak/aoc2020/lib"
)

var input = []*lib.Range{
	&lib.Range{45, 50, ""},
	&lib.Range{33, 44, ""},
	&lib.Range{13, 40, ""},
	&lib.Range{5, 7, ""},
	&lib.Range{6, 8, ""},
	&lib.Range{1, 3, ""},
}

func TestFindMatch(t *testing.T) {
	tree, err := lib.NewIntervalTree(input)
	if err != nil {
		panic(err)
	}

	result := tree.Find(40)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, lib.Range{13, 40, ""}, *(result[0]))
	assert.Equal(t, lib.Range{33, 44, ""}, *(result[1]))

}

func TestNoMatch(t *testing.T) {
	tree, err := lib.NewIntervalTree(input)
	if err != nil {
		panic(err)
	}

	result := tree.Find(11)
	assert.Equal(t, 0, len(result))
}

func TestRangeFind(t *testing.T) {
	tree, err := lib.NewIntervalTree(input)
	assert.Nil(t, err)
	result := tree.FindRange(&lib.Range{4, 15, ""})
	assert.Equal(t, 3, len(result))
	expected := input[2:5]
	assert.Equal(t, expected, result)
}
