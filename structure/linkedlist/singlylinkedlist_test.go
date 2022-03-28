package linkedlist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	t.Parallel()
	list := NewSinglyLinkedList()
	require.Equal(t, 0, list.Count())
	require.Equal(t, true, list.Empty())

	for i := 0; i < 5; i++ {
		list.PushFront(i)
	}

	require.Equal(t, 5, list.Count())
	require.Equal(t, false, list.Empty())

	ints := make([]int, 0, 5)
	for {
		val := list.Front()
		if val != nil {
			ints = append(ints, val.(int))
		} else {
			break
		}
	}
	require.Equal(t, []int{4, 3, 2, 1, 0}, ints)
	require.Equal(t, 0, list.Count())
	require.Equal(t, true, list.Empty())

	list.Clear()
	require.Equal(t, true, list.Empty())

	for i := 0; i < 5; i++ {
		list.PushBack(i)
	}

	require.Equal(t, 5, list.Count())
	require.Equal(t, false, list.Empty())

	ints2 := make([]int, 0, 5)
	for {
		val := list.Back()
		if val != nil {
			ints2 = append(ints2, val.(int))
		} else {
			break
		}
	}
	require.Equal(t, []int{4, 3, 2, 1, 0}, ints2)
	require.Equal(t, 0, list.Count())
	require.Equal(t, true, list.Empty())

	for i := 0; i < 5; i++ {
		list.PushFront(i)
	}

	list.Reverse()
	ints3 := make([]int, 0, 5)
	for {
		val := list.Front()
		if val != nil {
			ints3 = append(ints3, val.(int))
		} else {
			break
		}
	}
	require.Equal(t, []int{0, 1, 2, 3, 4}, ints3)
	list.Clear()
}
