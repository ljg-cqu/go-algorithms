package queue

type Queue interface {
	Enqueue(e any)
	Dequeue() any
	Peek() any
	Count() int
	Empty() bool
	Clear()
}
