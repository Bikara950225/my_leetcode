package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		src func() []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				src: func() []int {
					n := 100000
					ret := make([]int, 0, n)
					for ; n >= 0; n-- {
						ret = append(ret, n)
					}
					return ret
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srcList := tt.args.src()
			HeapSort(srcList)

			for i := 0; i < len(srcList)-1; i++ {
				if srcList[i] > srcList[i+1] {
					t.Errorf("TestHeapSort() error; not asc: \n%+v", srcList)
					return
				}
			}
		})
	}
}
