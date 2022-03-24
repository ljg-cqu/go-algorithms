package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCap(t *testing.T) {
	Push("hello world")
	e := Peak().(string)
	require.Equal(t, "hello world", e)
	require.Equal(t, Len(), 1)
	require.Equal(t, Cap(), 64)
	require.True(t, !Empty())

	eP := Pop().(string)
	require.Equal(t, "hello world", eP)
	require.Equal(t, Len(), 0)
	require.Equal(t, Cap(), 64)
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
	require.Equal(t, 1, Len())

	Clear()
}
