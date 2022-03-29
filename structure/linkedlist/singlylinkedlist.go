package linkedlist

import (
	"github.com/ljg-cqu/gotypes"
)

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
	////approach 1:
	//if s.nodes < 2 {
	//	return
	//}
	//
	//first := s.head.Next
	//
	//var hHead = new(gotypes.Node)
	//var hTail = first
	//for first != nil {
	//	second := first.Next
	//
	//	first.Next = hHead.Next
	//	hHead.Next = first
	//
	//	first = second
	//}
	//
	//s.head.Next = hHead.Next
	//s.tail = hTail
	//
	//if s.nodes < 2 {
	//	return
	//}

	//// approach 2:
	//first := s.head.Next
	//s.tail = first
	//s.head.Next = nil
	//s.reverseRec(first)

	// approach 3:
	first := s.head.Next
	s.tail = first
	s.head.Next = nil

	for first != nil {
		second := first.Next
		first.Next = s.head.Next
		s.head.Next = first
		first = second
	}
}

func (s *singlyLinkedList) reverseRec(first *gotypes.Node) {
	if first == nil {
		return
	}
	second := first.Next
	first.Next = s.head.Next
	s.head.Next = first
	s.reverseRec(second)
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
