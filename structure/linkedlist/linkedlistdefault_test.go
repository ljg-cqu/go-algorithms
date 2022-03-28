package linkedlist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultLinkedList(t *testing.T) {
	t.Parallel()

	require.Equal(t, 0, Count())
	require.Equal(t, true, Empty())

	for i := 0; i < 5; i++ {
		PushFront(i)
	}

	require.Equal(t, 5, Count())
	require.Equal(t, false, Empty())

	ints := make([]int, 0, 5)
	for {
		val := Front()
		if val != nil {
			ints = append(ints, val.(int))
		} else {
			break
		}
	}
	require.Equal(t, []int{4, 3, 2, 1, 0}, ints)
	require.Equal(t, 0, Count())
	require.Equal(t, true, Empty())

	Clear()
	require.Equal(t, true, Empty())

	for i := 0; i < 5; i++ {
		PushBack(i)
	}

	require.Equal(t, 5, Count())
	require.Equal(t, false, Empty())

	ints2 := make([]int, 0, 5)
	for {
		val := Back()
		if val != nil {
			ints2 = append(ints2, val.(int))
		} else {
			break
		}
	}
	require.Equal(t, []int{4, 3, 2, 1, 0}, ints2)
	require.Equal(t, 0, Count())
	require.Equal(t, true, Empty())

	for i := 0; i < 5; i++ {
		PushFront(i)
	}

	Reverse()
	ints3 := make([]int, 0, 5)
	for {
		val := Front()
		if val != nil {
			ints3 = append(ints3, val.(int))
		} else {
			break
		}
	}
	require.Equal(t, []int{0, 1, 2, 3, 4}, ints3)
	Clear()
}
