package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupDeque() Deque {
	return NewDeque([]int{
		9, 4, 7, 3, 2, 6,
	})
}

func TestNewDeque(t *testing.T) {
	d := setupDeque()
	head, ok := d.PeekTop()
	assert.Equal(t, 9, head)
	assert.True(t, ok)

	tail, ok := d.PeekBottom()
	assert.Equal(t, 6, tail)
	assert.True(t, ok)

	assert.Equal(t, 6, d.Count())

	assert.Equal(t, 9, d.PopTop())
	assert.Equal(t, 4, d.PopTop())
	assert.Equal(t, 7, d.PopTop())
	assert.Equal(t, 3, d.PopTop())
	assert.Equal(t, 2, d.PopTop())
	assert.Equal(t, 6, d.PopTop())
	assert.Equal(t, 0, d.PopTop())
	assert.True(t, d.IsEmpty())
	assert.Equal(t, 0, d.Count())
}

func TestMoveTopToBottom(t *testing.T) {
	d := setupDeque()
	r := d.PopTop()
	d.PushBottom(r)

	head, ok := d.PeekTop()
	assert.Equal(t, 4, head)
	assert.True(t, ok)

	tails, ok := d.PeekBottom()
	assert.Equal(t, 9, tails)
	assert.True(t, ok)
}

func TestEmpty(t *testing.T) {
	d := NewDeque([]int{})

	_, ok := d.PeekTop()
	assert.False(t, ok)
	_, ok = d.PeekBottom()
	assert.False(t, ok)

	n := d.PopTop()
	assert.Equal(t, 0, n)
	n = d.PopBottom()
	assert.Equal(t, 0, n)

	assert.True(t, d.IsEmpty())
}

func TestPushTop(t *testing.T) {
	d := NewDeque([]int{})
	d.PushTop(7)

	n, ok := d.PeekBottom()
	assert.True(t, ok)
	assert.Equal(t, 7, n)

	d.PushTop(9)
	n, ok = d.PeekBottom()
	assert.True(t, ok)
	assert.Equal(t, 7, n)

	n, ok = d.PeekTop()
	assert.True(t, ok)
	assert.Equal(t, 9, n)

	assert.False(t, d.IsEmpty())
}

func TestPopBottom(t *testing.T) {
	d := setupDeque()

	assert.Equal(t, 6, d.PopBottom())
	assert.Equal(t, 2, d.PopBottom())
	assert.Equal(t, 3, d.PopBottom())
	assert.Equal(t, 7, d.PopBottom())
	assert.Equal(t, 4, d.PopBottom())
	assert.Equal(t, 9, d.PopBottom())
	assert.Equal(t, 0, d.PopBottom())
	assert.True(t, d.IsEmpty())
}

func TestTakeTop(t *testing.T) {
	t.Run("take subset", func(t *testing.T) {
		d := setupDeque()
		o := d.TakeTop(3)

		assert.Equal(t, 6, d.Count())
		assert.Equal(t, 3, o.Count())

		n, ok := o.PeekBottom()
		assert.Equal(t, 7, n)
		assert.True(t, ok)

		n = o.PopBottom()
		assert.Equal(t, 7, n)

		// ensure base deque is unmodified
		assert.Equal(t, 6, d.Count())
	})

	t.Run("try taking superset", func(t *testing.T) {
		d := setupDeque()
		o := d.TakeTop(100)

		assert.Equal(t, o.Count(), 6)

		n, ok := o.PeekBottom()
		assert.Equal(t, 6, n)
		assert.True(t, ok)

		n = o.PopBottom()
		assert.Equal(t, 6, n)

		// ensure base deque is unmodified
		assert.Equal(t, 6, d.Count())
	})

	t.Run("try using negative N", func(t *testing.T) {
		d := setupDeque()
		o := d.TakeTop(-1)

		assert.True(t, o.IsEmpty())

		o = d.TakeTop(0)
		assert.True(t, o.IsEmpty())
	})
}

func TestVisit(t *testing.T) {
	t.Run("visit on empty deque", func(t *testing.T) {
		f := func(int) {
			assert.FailNow(t, "wasn't supposed to get called!")
		}

		d := NewDeque([]int{})
		d.Visit(f)
	})

	t.Run("visit a normal deque", func(t *testing.T) {
		d := setupDeque()
		i := 0
		vals := make([]int, d.Count())

		f := func(n int) {
			vals[i] = n
			i++
		}

		d.Visit(f)

		assert.Equal(t, []int{
			9, 4, 7, 3, 2, 6,
		}, vals)

	})
}
