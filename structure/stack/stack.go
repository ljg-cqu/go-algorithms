package stack

type Stack interface {
	Push(e any)
	Pop() any
	Peak() any
	Count() int
	Empty() bool
	Clear()
}
