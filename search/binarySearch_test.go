package search

import (
	"github.com/juju/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()
	type testT struct {
		name           string
		inputArr       []int
		inputTarget    int
		inputLowIndex  int
		inputHighIndex int
		want           int
		wantErr        error
	}

	var scenarios = []testT{
		{"error path: nil arr", []int(nil), 1, 0, 1, -1, errors.NotFound},
		{"error path: no elem", []int{}, 1, -1, 1, -1, errors.NotFound},
		{"error path: not found", []int{1, 2, 3, 4, 5}, 6, 0, 4, -1, errors.NotFound},
		{"error path: low index out bound", []int{1, 2, 3, 4, 5}, 1, -1, 4, -1, errors.NotValid},
		{"error path: high index out bound", []int{1, 2, 3, 4, 5}, 1, 0, 5, -1, errors.NotValid},
		{"error path: high index out bound", []int{1, 2, 3, 4, 5}, 1, 0, 5, -1, errors.NotValid},
		{"happy path: index first", []int{1, 2, 3, 4, 5}, 1, 0, 4, 0, nil},
		{"happy path: index last", []int{1, 2, 3, 4, 5}, 5, 0, 4, 4, nil},
	}

	for _, c := range scenarios {
		c := c // capture range variable
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			gotIndex, gotErr := BinarySearch(c.inputArr, c.inputTarget, c.inputLowIndex, c.inputHighIndex)
			require.True(t, gotIndex == c.want && errors.Is(gotErr, c.wantErr))
			gotIndex, gotErr = BinarySearchRec(c.inputArr, c.inputTarget, c.inputLowIndex, c.inputHighIndex)
			require.True(t, gotIndex == c.want && errors.Is(gotErr, c.wantErr))
		})
	}
}
