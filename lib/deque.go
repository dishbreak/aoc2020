package lib

type Deque interface {
	PushTop(data int)
	PeekTop() (int, bool)
	PopTop() int
	PushBottom(data int)
	PeekBottom() (int, bool)
	PopBottom() int
	IsEmpty() bool
}

type deque struct {
	top    *dequeNode
	bottom *dequeNode
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
}
