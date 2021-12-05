package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoopSize(t *testing.T) {
	loopSize := getLoopSize(5764801)
	assert.Equal(t, 8, loopSize)
}

func TestGetEncKey(t *testing.T) {
	assert.Equal(t, 14897079, getEncKey(17807724, 8))
	assert.Equal(t, 14897079, getEncKey(5764801, 11))
}
