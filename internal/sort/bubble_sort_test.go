package sort

import "testing"

func TestBubbleSort(t *testing.T) {
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
					n := 1000
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
			BubbleSort(srcList)

			for i := 0; i < len(srcList)-1; i++ {
				if srcList[i] > srcList[i+1] {
					t.Errorf("TestBubbleSort() not asc: \n%+v", srcList)
					return
				}
			}
		})
	}
}
