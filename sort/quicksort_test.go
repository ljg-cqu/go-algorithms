package sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// todo: fuzz

func TestQuickSort(t *testing.T) {
	t.Parallel()
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"quick sort", args{[]int{5, 4, 3, 0, 10}}, []int{0, 3, 4, 5, 10}},
		{"quick sort", args{[]int{105, 44, 30, 15, 2}}, []int{2, 15, 30, 44, 105}},
		{"quick sort:empty input", args{[]int{}}, []int{}},
		{"quick sort:nil input", args{nil}, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			QuickSort(tt.args.arr)
			require.Equal(t, tt.want, tt.args.arr)

		})
	}
}
