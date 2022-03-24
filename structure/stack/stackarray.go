package stack

type stackArray []any

func NewStackArray() Stack {
	stack := make(stackArray, 0, 64)
	return &stack
}

func (s *stackArray) Push(e any) {
	*s = append(*s, e)
}

func (s *stackArray) Pop() any {
	if len(*s) == 0 {
		return nil
	}

	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e
}

func (s *stackArray) Peak() any {
	if s.Empty() {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *stackArray) Len() int {
	return len(*s)
}

func (s *stackArray) Cap() int {
	return cap(*s)
}

func (s *stackArray) Empty() bool {
	return len(*s) == 0
}

func (s *stackArray) Clear() {
	if !s.Empty() {
		*s = make(stackArray, 0, 64)
	}
}
