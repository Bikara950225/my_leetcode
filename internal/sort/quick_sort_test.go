package sort

import (
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		src func() []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				src: func() []int {
					var n int64 = 10000
					ret := make([]int, 0, n)
					for ; n >= 0; n-- {
						ret = append(ret, int(n))
					}
					return ret
				},
			},
		}, {
			name: "test2",
			args: args{
				src: func() []int {
					return []int{-1, 0, 1, 2, -1, -4}
				},
			},
		},
		{
			name: "test3",
			args: args{
				src: func() []int {
					return []int{2, 2, 2, 1, 1}
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src := tt.args.src()
			QuickSort(src)

			if !slices.IsSortedFunc(src, func(i, j int) int {
				return i - j
			}) {
				t.Errorf("TestQuickSort() error, not ase: \n%+v", src)
				return
			}
		})
	}
}
