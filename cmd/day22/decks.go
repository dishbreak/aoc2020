package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
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
		acc += (n * mult)
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

type recursiveCombatGame struct {
	*combatGame
	pastRounds map[string]int
}

func buildRecursiveCombatGame(input [][]string) *recursiveCombatGame {
	return &recursiveCombatGame{
		combatGame: buildGame(input),
		pastRounds: make(map[string]int),
	}
}

func (r *recursiveCombatGame) hashRound() string {
	sha256 := sha256.New()
	hasher := func(n int) {
		sha256.Write([]byte(strconv.Itoa(n)))
	}

	for _, deque := range []lib.Deque{r.player1, r.player2} {
		deque.Visit(hasher)
	}

	return base64.URLEncoding.EncodeToString(sha256.Sum(nil))
}

func (r *recursiveCombatGame) playGame() (p1score, p2score int) {
	p1score, p2score = 0, 0
	p1wins, p2wins := false, false
	for ; !(p1wins || p2wins); p1wins, p2wins = r.playRound() {
	}

	fmt.Println("Post game results")
	fmt.Printf("player 1 deck: %s\n", r.player1)
	fmt.Printf("player 2 deck: %s\n\n", r.player2)

	if p1wins {
		p1score = scoreDeck(r.player1)
	} else {
		p2score = scoreDeck(r.player2)
	}
	return
}

func (r *recursiveCombatGame) playRound() (player1wins, player2wins bool) {
	player1wins = false
	player2wins = false

	fmt.Printf("=======\n\n")
	if r.player1.IsEmpty() {
		fmt.Println("player 2 wins!")
		player2wins = true
		return
	}

	if r.player2.IsEmpty() {
		fmt.Println("player 1 wins!")
		player1wins = true
		return
	}

	roundHash := r.hashRound()
	if _, ok := r.pastRounds[roundHash]; ok {
		fmt.Println("infinite game protection! player 1 wins!")
		player1wins = true
		return
	}
	r.pastRounds[roundHash]++

	fmt.Printf("Player 1 deck: %s\n", r.player1)
	fmt.Printf("player 2 deck: %s\n", r.player2)
	fmt.Println("")
	p1Card := r.player1.PopTop()
	p2Card := r.player2.PopTop()

	fmt.Printf("player 1 played: %d\n", p1Card)
	fmt.Printf("player 2 played: %d\n", p2Card)
	fmt.Println("")

	if r.player1.Count() >= p1Card && r.player2.Count() >= p2Card {
		fmt.Println("Playing a subgame to determine the winner.")
		fmt.Printf("\n\n")
		subgame := &recursiveCombatGame{
			combatGame: &combatGame{
				player1: r.player1.TakeTop(p1Card),
				player2: r.player2.TakeTop(p2Card),
			},
			pastRounds: make(map[string]int),
		}
		sgP1, _ := subgame.playGame()
		if sgP1 > 0 {
			fmt.Println("player 1 wins the subgame")
			r.player1.PushBottom(p1Card)
			r.player1.PushBottom(p2Card)
			return
		}
		fmt.Println("player 2 wins the subgame.")
		r.player2.PushBottom(p2Card)
		r.player2.PushBottom(p1Card)
		return
	}

	if p1Card > p2Card {
		fmt.Println("player 1 wins the round")
		r.player1.PushBottom(p1Card)
		r.player1.PushBottom(p2Card)
		return
	}
	fmt.Println("player 2 wins the round")
	r.player2.PushBottom(p2Card)
	r.player2.PushBottom(p1Card)
	return
}
