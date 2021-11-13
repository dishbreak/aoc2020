package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayRound(t *testing.T) {
	g := newShellGame("389125467")

	for i := 0; i < 10; i++ {
		g.playRound()
	}

	assert.Equal(t, "92658374", g.String())
}
