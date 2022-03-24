package stack

import "sync"

var defaultStack Stack

var once = &sync.Once{}

func initDefaultStack() {
	once.Do(func() {
		defaultStack = NewStackArray()
	})
}

func Push(e any) {
	initDefaultStack()
	defaultStack.Push(e)
}

func Pop() any {
	initDefaultStack()
	return defaultStack.Pop()
}

func Peak() any {
	initDefaultStack()
	return defaultStack.Peak()
}

func Len() int {
	initDefaultStack()
	return defaultStack.Len()
}

func Cap() int {
	initDefaultStack()
	return defaultStack.Cap()
}

func Empty() bool {
	initDefaultStack()
	return defaultStack.Empty()
}

func Clear() {
	initDefaultStack()
	defaultStack.Clear()
}
