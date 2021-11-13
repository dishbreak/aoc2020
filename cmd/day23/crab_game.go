package main

import (
	"strconv"
	"strings"

	"github.com/dishbreak/aoc2020/lib"
)

type bigShellGame struct {
	circle *lib.LinkedList
	cups   map[int]*lib.LinkedListNode
	iter   *lib.LinkedListNode
}

const (
	numCups = 1000000
)

func newCrabShellGame(input string) *bigShellGame {
	g := &bigShellGame{
		cups: make(map[int]*lib.LinkedListNode),
	}

	b := &lib.LinkedListBuilder{}

	max := -1

	for _, char := range strings.Split(input, "") {
		parsed, _ := strconv.Atoi(char)
		if parsed > max {
			max = parsed
		}
		n := b.AddItem(parsed)
		g.cups[parsed] = n
	}

	for i := max + 1; i <= numCups; i++ {
		n := b.AddItem(i)
		g.cups[i] = n
	}

	g.circle = b.GetList()

	// make this a circular list.
	g.circle.Tail.Next = g.circle.Head

	//start with the first supplied cup
	g.iter = g.circle.Head

	return g
}

func wraparound(i int) int {
	if i == 0 {
		return numCups
	}
	return i
}

func (g *bigShellGame) playRound() {
	// select the cups to remove
	cupStart := g.iter.Next
	cupEnd := cupStart.Next.Next

	// remove the cups for simplicity's sake.
	g.iter.Next = cupEnd.Next
	cupEnd.Next = nil

	// register all the labels removed.
	removedCups := make(map[int]*lib.LinkedListNode)
	for iter := cupStart; iter != nil; iter = iter.Next {
		removedCups[iter.Data] = iter
	}

	// determine the target
	target := wraparound(g.iter.Data - 1)

	// decrement the target until it is no longer one of the removed cups.
	for _, ok := removedCups[target]; ok; _, ok = removedCups[target] {
		target = wraparound(target - 1)
	}

	// lookup the cup with the target label
	destCup := g.cups[target]

	// insert the cups removed following the destination cup.
	cupEnd.Next = destCup.Next
	destCup.Next = cupStart

	// move the current cup over.
	g.iter = g.iter.Next
}

type gameCup *lib.LinkedListNode

func (g *bigShellGame) getCup(label int) gameCup {
	return gameCup(g.cups[label])
}
