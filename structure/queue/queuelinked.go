package queue

type node struct {
	elem any
	next *node
}

type queueLinked struct {
	nodes int
	head  *node
	tail  *node
}

func NewQueueLinked() Queue {
	return &queueLinked{0, new(node), nil}
}

func (q *queueLinked) Enqueue(e any) {
	newNode := &node{e, nil}
	defer func() { q.nodes++ }()

	if q.tail == nil {
		q.head.next = newNode
		q.tail = q.head.next
		return
	}
	q.tail.next = newNode
	q.tail = q.tail.next
}

func (q *queueLinked) Dequeue() any {
	if q.nodes == 0 {
		return nil
	}
	defer func() { q.nodes-- }()

	e := q.head.next.elem
	q.head.next = q.head.next.next
	return e
}

func (q *queueLinked) Peek() any {
	if q.nodes == 0 {
		return nil
	}
	return q.tail.elem
}

func (q *queueLinked) Count() int {
	return q.nodes
}

func (q *queueLinked) Empty() bool {
	return q.nodes == 0
}

func (q *queueLinked) Clear() {
	q.head.next = nil
	q.tail = nil
	q.nodes = 0
}
