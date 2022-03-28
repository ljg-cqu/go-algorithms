package linkedlist

import "github.com/ljg-cqu/gotypes"

type doublyLinkedList struct {
	nodes int
	head  *gotypes.Node
	tail  *gotypes.Node
}

func NewDoublyLinkedList() gotypes.LinkedList {
	return &doublyLinkedList{head: new(gotypes.Node)}
}

func (s *doublyLinkedList) PushFront(val any) {
	defer func() { s.nodes++ }()

	newNode := &gotypes.Node{Val: val}

	if s.nodes == 0 {
		s.head.Next = newNode
		newNode.Prev = s.head
		s.tail = s.head.Next
		return
	}

	newNode.Next = s.head.Next
	newNode.Next.Prev = newNode

	s.head.Next = newNode
	newNode.Prev = s.head
}

func (s *doublyLinkedList) PushBack(val any) {
	defer func() { s.nodes++ }()

	newNode := &gotypes.Node{Val: val}

	if s.nodes == 0 {
		s.head.Next = newNode
		newNode.Prev = s.head
		s.tail = s.head.Next

		return
	}

	s.tail.Next = newNode
	newNode.Prev = s.tail

	s.tail = newNode
}

func (s *doublyLinkedList) Front() any {
	if s.nodes == 0 {
		return nil
	}

	node := s.head.Next
	s.head.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = s.head
	}
	s.nodes--
	return node.Val
}

func (s *doublyLinkedList) Back() any {
	if s.nodes == 0 {
		return nil
	}

	val := s.tail.Val

	s.tail = s.tail.Prev
	s.tail.Next.Prev = nil
	s.tail.Next = nil

	s.nodes--

	return val
}

func (s *doublyLinkedList) Reverse() {
	if s.nodes <= 1 {
		return
	}

	var hHead = new(gotypes.Node)
	for first := s.head.Next; first != nil; {
		second := first.Next

		first.Next = hHead.Next
		if hHead.Next != nil {
			hHead.Next.Prev = first
		}

		hHead.Next = first
		first.Prev = hHead

		first = second
	}

	s.head.Next = hHead.Next
	hHead.Next.Prev = s.head
}

func (s *doublyLinkedList) Count() int {
	return s.nodes
}

func (s *doublyLinkedList) Empty() bool {
	return s.nodes == 0
}

func (s *doublyLinkedList) Clear() {
	s.nodes = 0
	s.head.Next = nil
	s.tail = nil
}
