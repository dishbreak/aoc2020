package lib

// StringStack is a stack that...well, it holds strings.
type StringStack []string

// Push adds a new item to the top of the Stack
func (s *StringStack) Push(input string) {
	*s = append(*s, input)
}

// IsEmpty returns true when the stack has no items, false otherwise.
func (s *StringStack) IsEmpty() bool {
	return len(*s) == 0
}

// Pop removes the most-recently added item from the stack and returns it.
// If the stack is empty, we'll return an empty string and a false boolean.
func (s *StringStack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}

	index := len(*s) - 1
	result := (*s)[index]
	*s = (*s)[:index]
	return result, true
}

// Peek will let us see the most-recently added item without removing it from
// the stack.
func (s *StringStack) Peek() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	return (*s)[len(*s)-1], true
}
