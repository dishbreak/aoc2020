package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildLinkedList(t *testing.T) {
	b := &LinkedListBuilder{}

	b.AddItem(5)
	b.AddItem(7)
	end := b.AddItem(3)

	l := b.GetList()

	assert.Equal(t, "5 -> 7 -> 3 -> nil", l.String())
	assert.Nil(t, end.Next)

}
