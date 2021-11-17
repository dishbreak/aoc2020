package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoopSize(t *testing.T) {
	loopSize, err := getLoopSize(5764801)
	assert.Equal(t, 8, loopSize)
	assert.Nil(t, err)
}

func TestGetEncKey(t *testing.T) {
	assert.Equal(t, 5764801, getEncKey(1, 8))
	assert.Equal(t, 14897079, getEncKey(17807724, 8))
}
