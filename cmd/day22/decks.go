package main

import (
	"strconv"

	"github.com/dishbreak/aoc2020/lib"
)

type combatGame struct {
	player1 lib.Deque
	player2 lib.Deque
}

func buildGame(input [][]string) *combatGame {
	return &combatGame{
		player1: buildDeck(input[0]),
		player2: buildDeck(input[1]),
	}
}

func buildDeck(input []string) lib.Deque {
	cards := make([]int, len(input)-1)
	for idx, str := range input[1:] {
		parsed, _ := strconv.Atoi(str)
		cards[idx] = parsed
	}

	return lib.NewDeque(cards)
}

func scoreDeck(d lib.Deque) int {
	acc := 0
	mult := d.Count()

	tally := func(n int) {
		acc += n * mult
		mult--
	}

	d.Visit(tally)

	return acc
}

func (c *combatGame) playRound() bool {
	if c.player1.IsEmpty() || c.player2.IsEmpty() {
		return false
	}

	p1 := c.player1.PopTop()
	p2 := c.player2.PopTop()

	if p1 > p2 {
		c.player1.PushBottom(p1)
		c.player1.PushBottom(p2)
	} else {
		c.player2.PushBottom(p2)
		c.player2.PushBottom(p1)
	}

	return true
}

func (c *combatGame) scoreGame() int {
	p1 := scoreDeck(c.player1)
	p2 := scoreDeck(c.player2)

	if p1 > p2 {
		return p1
	}

	return p2
}
