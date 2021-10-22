package main

import (
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	p       [][]int
	seen    map[string]int
	recurse bool
}

func newGame(input [][]string) (g *game) {
	g = &game{
		seen: make(map[string]int),
	}
	g.p = make([][]int, 2)

	for i, player := range input {
		cards := player[1:]
		g.p[i] = make([]int, len(cards))
		for j, card := range cards {
			parsed, _ := strconv.Atoi(card)
			g.p[i][j] = parsed
		}
	}

	return
}

func scoreDeck(d []int) (acc int) {
	acc = 0

	for idx, c := range d {
		acc += (len(d) - idx) * c
	}

	return
}

func (g *game) winningScore() int {
	for _, p := range g.p {
		if len(p) > 0 {
			return scoreDeck(p)
		}
	}
	return 0
}

func deckAsStr(d []int) string {
	j := make([]string, len(d))
	for idx, c := range d {
		j[idx] = strconv.Itoa(c)
	}

	return fmt.Sprintf("[%s]", strings.Join(j, ", "))
}

func (g *game) playRound() int {
	key := deckAsStr(g.p[0]) + deckAsStr(g.p[1])
	if _, ok := g.seen[key]; ok {
		return 1
	}
	g.seen[key] = 1

	p1 := g.p[0][0]
	p2 := g.p[1][0]

	g.p[0] = g.p[0][1:]
	g.p[1] = g.p[1][1:]

	p1wins := p1 > p2
	if g.recurse && len(g.p[0]) >= p1 && len(g.p[1]) >= p2 {
		sub := &game{
			p:       make([][]int, 2),
			seen:    make(map[string]int),
			recurse: true,
		}
		sub.p[0] = make([]int, p1)
		sub.p[1] = make([]int, p2)
		copy(sub.p[0], g.p[0][:p1])
		copy(sub.p[1], g.p[1][:p2])

		p1wins = sub.playGame() == 1
	}

	if p1wins {
		g.p[0] = append(g.p[0], p1, p2)
	} else {
		g.p[1] = append(g.p[1], p2, p1)
	}

	if len(g.p[0]) == 0 {
		return 2
	} else if len(g.p[1]) == 0 {
		return 1
	}

	return 0
}

func (g *game) playGame() int {
	w := 0
	for ; w == 0; w = g.playRound() {

	}
	return w
}
