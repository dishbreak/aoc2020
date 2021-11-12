package lib

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
