package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStackDefault(t *testing.T) {
	Push("hello world")
	e := Peak().(string)
	require.Equal(t, "hello world", e)
	require.Equal(t, Count(), 1)
	require.True(t, !Empty())

	eP := Pop().(string)
	require.Equal(t, "hello world", eP)
	require.Equal(t, Count(), 0)
	require.True(t, Empty())

	ePP := Pop()
	require.Equal(t, nil, ePP)

	Push(e)
	Clear()
	require.True(t, Empty())

	Push(e)
	Push(100)
	eN := Pop().(int)
	require.Equal(t, 100, eN)
	require.Equal(t, 1, Count())

	Clear()
	for i := 0; i < 5; i++ {
		Push(i)
	}
	require.Equal(t, 5, Count())

	elems := make([]int, 0, 5)
	for {
		e := Pop()
		if e == nil {
			break
		}
		elems = append(elems, e.(int))
	}
	require.Equal(t, []int{4, 3, 2, 1, 0}, elems)
}
