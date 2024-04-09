package code_mianshi

import (
	"reflect"
	"testing"
)

func Test_findNumber(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{
			name: "case1",
			args: args{
				src: []int{1, 2, 4, 3, 10, 13, 15, 12, 16, 18, 17},
			},
			wantRet: []int{
				1, 2, 10, 16,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := findNumber(tt.args.src); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("findNumber() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
