package lib

import (
	"strconv"
	"strings"
)

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
}

type LinkedListBuilder struct {
	list *LinkedList
}

func (l *LinkedList) String() string {
	b := strings.Builder{}

	for iter := l.Head; iter != nil; iter = iter.Next {
		b.WriteString(strconv.Itoa(iter.Data))
		b.WriteString(" -> ")
	}

	b.WriteString("nil")

	return b.String()
}

func (b *LinkedListBuilder) AddItem(data int) *LinkedListNode {

	n := &LinkedListNode{
		Data: data,
	}

	if b.list == nil {
		b.list = &LinkedList{
			Head: n,
			Tail: n,
		}
	} else {
		b.list.Tail.Next = n
		b.list.Tail = n
	}

	return n
}

func (b *LinkedListBuilder) GetList() *LinkedList {
	return b.list
}
