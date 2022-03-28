package linkedlist

import "github.com/ljg-cqu/gotypes"

type singlyLinkedList struct {
	nodes int
	head  *gotypes.Node
	tail  *gotypes.Node
}

func NewSinglyLinkedList() gotypes.LinkedList {
	return &singlyLinkedList{head: new(gotypes.Node)}
}

func (s *singlyLinkedList) PushFront(val any) {
	newNode := gotypes.Node{Val: val, Next: s.head.Next}
	s.head.Next = &newNode

	if s.nodes == 0 {
		s.tail = s.head.Next
	}
	s.nodes++
}

func (s *singlyLinkedList) Back() any {
	if s.nodes == 0 {
		return nil
	}

	if s.nodes == 1 {
		val := s.head.Next.Val
		s.head.Next = nil
		s.tail = nil
		s.nodes--
		return val
	}

	p := s.head.Next
	for ; p.Next != s.tail; p = p.Next {
	}
	e := s.tail.Val
	p.Next = nil
	s.tail = p
	s.nodes--
	return e
}

func (s *singlyLinkedList) Front() any {
	if s.nodes == 0 {
		return nil
	}

	val := s.head.Next.Val
	s.head.Next = s.head.Next.Next
	s.nodes--
	return val
}

func (s *singlyLinkedList) PushBack(val any) {
	defer func() { s.nodes++ }()

	newE := &gotypes.Node{Val: val}
	if s.nodes == 0 {
		s.head.Next = newE
		s.tail = s.head.Next
		return
	}

	s.tail.Next = newE
	s.tail = s.tail.Next
}

func (s *singlyLinkedList) Reverse() {
	var hHead = new(gotypes.Node)
	for first := s.head.Next; first != nil; {
		second := first.Next

		first.Next = hHead.Next
		hHead.Next = first

		first = second
	}

	s.head.Next = hHead.Next
}

func (s *singlyLinkedList) Count() int {
	return s.nodes
}

func (s *singlyLinkedList) Empty() bool {
	return s.nodes == 0
}

func (s *singlyLinkedList) Clear() {
	s.nodes = 0
	s.head.Next = nil
	s.tail = nil
}
