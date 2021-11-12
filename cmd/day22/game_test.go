package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := newGame(input)

	assert.Equal(t, []int{9, 2, 6, 3, 1}, g.p[0])
}

func TestScoreDeck(t *testing.T) {
	assert.Equal(t, 306, scoreDeck([]int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1}))
}

func TestWinningScore(t *testing.T) {
	g := &game{
		p: [][]int{
			{},
			{3, 2, 10, 6, 8, 5, 9, 4, 7, 1},
		},
	}

	assert.Equal(t, 306, g.winningScore())
}
