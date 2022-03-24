package stack

type Stack interface {
	Push(e any)
	Pop() any
	Peak() any
	Len() int
	Cap() int
	Empty() bool
	Clear()
}
