package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type Deque interface {
	PushTop(data int)
	PeekTop() (int, bool)
	PopTop() int
	PushBottom(data int)
	PeekBottom() (int, bool)
	PopBottom() int
	IsEmpty() bool
	Count() int
	TakeTop(int) Deque
	Visit(NodeVisitor)
	String() string
}

type deque struct {
	top    *dequeNode
	bottom *dequeNode
	count  int
}

type dequeNode struct {
	data   int
	top    *dequeNode
	bottom *dequeNode
}

func NewDeque(input []int) Deque {
	d := &deque{}
	for _, data := range input {
		d.PushBottom(data)
	}
	return d
}

func (d *deque) PushTop(data int) {
	n := &dequeNode{
		data: data,
	}

	n.bottom = d.top
	if d.top != nil {
		d.top.top = n
	}
	d.top = n

	if d.bottom == nil {
		d.bottom = n
	}
	d.count++
}

func (d *deque) PeekTop() (int, bool) {
	result := 0
	if d.IsEmpty() {
		return result, false
	}

	result = d.top.data
	return result, true
}

func (d *deque) PopTop() int {
	result, ok := d.PeekTop()
	if !ok {
		return result
	}

	d.top = d.top.bottom
	if d.top == nil {
		d.bottom = nil
	} else {
		d.top.top = nil
	}
	d.count--
	return result
}

func (d *deque) IsEmpty() bool {
	return d.top == nil
}

func (d *deque) PeekBottom() (int, bool) {
	result := 0
	if d.bottom == nil {
		return result, false
	}
	return d.bottom.data, true
}

func (d *deque) PopBottom() int {
	result, ok := d.PeekBottom()
	if !ok {
		return result
	}

	d.bottom = d.bottom.top
	if d.bottom == nil {
		d.top = nil
	}
	d.count--
	return result
}

func (d *deque) PushBottom(data int) {
	n := &dequeNode{
		data: data,
	}

	if d.bottom != nil {
		d.bottom.bottom = n
	}
	n.top = d.bottom
	d.bottom = n

	if d.top == nil {
		d.top = n
	}
	d.count++
}

func (d *deque) Count() int {
	return d.count
}

func (d *deque) TakeTop(n int) Deque {
	if n > d.count {
		n = d.count
	}

	if n < 0 {
		n = 0
	}

	if n == 0 {
		return NewDeque([]int{})
	}

	values := make([]int, n)

	con := d.top
	for i := 0; i < n; i++ {
		values[i] = con.data
		con = con.bottom
	}

	return NewDeque(values)
}

type NodeVisitor func(int)

func (d *deque) Visit(n NodeVisitor) {
	for con := d.top; con != nil; con = con.bottom {
		n(con.data)
	}
}

func (d *deque) String() string {
	result := make([]string, d.count)
	i := 0
	f := func(n int) {
		result[i] = strconv.Itoa(n)
		i++
	}
	d.Visit(f)

	return fmt.Sprintf("[%s] (%d)", strings.Join(result, ", "), d.count)
}
