package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack(t *testing.T) {
	t.Parallel()

	type stackTest struct {
		name  string
		stack Stack
	}

	stackTests := []stackTest{
		{"array stack", NewStackArray()},
		{"linked stack", NewStackLinked()},
	}

	for _, st := range stackTests {
		st := st
		t.Run(st.name, func(t *testing.T) {
			t.Parallel()
			st.stack.Push("hello world")
			e := st.stack.Peak().(string)
			require.Equal(t, "hello world", e)
			require.Equal(t, st.stack.Count(), 1)
			require.True(t, !st.stack.Empty())

			eP := st.stack.Pop().(string)
			require.Equal(t, "hello world", eP)
			require.Equal(t, st.stack.Count(), 0)
			require.True(t, st.stack.Empty())

			ePP := st.stack.Pop()
			require.Equal(t, nil, ePP)

			st.stack.Push(e)
			st.stack.Clear()
			require.True(t, Empty())

			st.stack.Push(e)
			st.stack.Push(100)
			eN := st.stack.Pop().(int)
			require.Equal(t, 100, eN)
			require.Equal(t, 1, st.stack.Count())

			st.stack.Clear()
			for i := 0; i < 5; i++ {
				st.stack.Push(i)
			}
			require.Equal(t, 5, st.stack.Count())

			elems := make([]int, 0, 5)
			for {
				e := st.stack.Pop()
				if e == nil {
					break
				}
				elems = append(elems, e.(int))
			}
			require.Equal(t, []int{4, 3, 2, 1, 0}, elems)
		})
	}
}
