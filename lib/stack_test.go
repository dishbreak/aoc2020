package lib_test

import (
	"testing"

	"github.com/dishbreak/aoc2020/lib"
	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	var s lib.StringStack

	s.Push("foo")
	s.Push("bar")
	s.Push("baz")

	expected := []string{
		"baz",
		"bar",
		"foo",
	}

	assert.False(t, s.IsEmpty())

	for _, result := range expected {
		actual, ok := s.Pop()
		assert.True(t, ok, "Stack should be non-empty")
		assert.Equal(t, result, actual, "Wrong value returned from stack.")
	}

	_, ok := s.Pop()
	assert.False(t, ok, "Stack should be empty")
}

func TestEmpty(t *testing.T) {
	var s lib.StringStack

	assert.True(t, s.IsEmpty())
}
