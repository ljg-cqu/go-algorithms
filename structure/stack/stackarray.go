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
	if len(*s) == 0 {
		return nil
	}
	return (*s)[len(*s)-1]
}

func (s *stackArray) Count() int {
	return len(*s)
}

func (s *stackArray) Empty() bool {
	return len(*s) == 0
}

func (s *stackArray) Clear() {
	*s = make(stackArray, 0, 64)
}
