package main

import (
	"testing"

	"github.com/dishbreak/aoc2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestBuildDeck(t *testing.T) {
	input := []string{"Player 1:", "3", "5", "7", "9", "11"}

	deck := buildDeck(input)
	assert.Equal(t, 5, deck.Count())

	n, ok := deck.PeekBottom()
	assert.Equal(t, 11, n)
	assert.True(t, ok)

	n, ok = deck.PeekTop()
	assert.Equal(t, 3, n)
	assert.True(t, ok)
}

func TestScoreDeck(t *testing.T) {
	input := []string{"Player 1:", "3", "5", "7", "9", "11"}

	deck := buildDeck(input)
	assert.Equal(t, 85, scoreDeck(deck))
}

func TestPlayRound(t *testing.T) {
	t.Run("player 1 wins round", func(t *testing.T) {
		inputs := [][]string{
			{"Player 1:", "3", "5", "7", "9", "11"},
			{"Player 2:", "4", "2", "6", "8", "10"},
		}

		game := buildGame(inputs)
		keepPlaying := game.playRound()
		assert.True(t, keepPlaying)

		n, ok := game.player1.PeekTop()
		assert.Equal(t, 5, n)
		assert.True(t, ok)
		n, ok = game.player2.PeekTop()
		assert.Equal(t, 2, n)
		assert.True(t, ok)
		n, ok = game.player1.PeekBottom()
		assert.Equal(t, 11, n)
		assert.True(t, ok)

		assert.Equal(t, 3, game.player2.PopBottom())
		assert.Equal(t, 4, game.player2.PopBottom())
	})
	t.Run("player 2 wins round", func(t *testing.T) {
		inputs := [][]string{
			{"Player 1:", "5", "7", "9", "11"},
			{"Player 2:", "2", "6", "8", "10"},
		}
		game := buildGame(inputs)

		keepPlaying := game.playRound()
		assert.True(t, keepPlaying)

		n, ok := game.player2.PeekTop()
		assert.Equal(t, 6, n)
		assert.True(t, ok)
		n, ok = game.player1.PeekTop()
		assert.Equal(t, 7, n)
		assert.True(t, ok)
		n, ok = game.player2.PeekBottom()
		assert.Equal(t, 10, n)
		assert.True(t, ok)

		assert.Equal(t, 2, game.player1.PopBottom())
		assert.Equal(t, 5, game.player1.PopBottom())
	})
	t.Run("game over", func(t *testing.T) {
		inputs := [][]string{
			{"Player 1:", "3", "5", "7", "9", "11"},
			{"Player 2:", "4", "2", "6", "8", "10"},
		}

		game := buildGame(inputs)
		game.player2 = lib.NewDeque([]int{})

		keepPlaying := game.playRound()
		assert.False(t, keepPlaying)
	})
}

func TestScoreGame(t *testing.T) {
	t.Run("player 1 wins", func(t *testing.T) {
		inputs := [][]string{
			{"Player 1:", "3", "5", "7", "9", "11"},
			{"Player 2:", "0"},
		}

		game := buildGame(inputs)

		assert.Equal(t, 85, game.scoreGame())
	})
	t.Run("player 2 wins", func(t *testing.T) {
		inputs := [][]string{
			{"Player 1:", "0"},
			{"Player 2:", "3", "5", "7", "9", "11"},
		}
		game := buildGame(inputs)

		assert.Equal(t, 85, game.scoreGame())
	})
}
