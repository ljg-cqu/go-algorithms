package stack

type node struct {
	elem any
	next *node
}

type stackLinked struct {
	nodes int
	head  *node
}

func NewStackLinked() Stack {
	return &stackLinked{
		nodes: 0,
		head:  new(node),
	}
}

func (s *stackLinked) Push(e any) {
	s.head.next = &node{e, s.head.next}
	s.nodes++
}

func (s *stackLinked) Pop() any {
	if s.nodes == 0 {
		return nil
	}
	e := s.head.next.elem
	s.head.next = s.head.next.next
	s.nodes--
	return e
}

func (s *stackLinked) Peak() any {
	if s.nodes == 0 {
		return nil
	}
	return s.head.next.elem
}

func (s *stackLinked) Count() int {
	return s.nodes
}

func (s *stackLinked) Empty() bool {
	return s.nodes == 0
}

func (s *stackLinked) Clear() {
	s.head.next = nil
	s.nodes = 0
}
