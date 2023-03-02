package monotonic_stack

import (
	"reflect"
	"testing"
)

func TestGetMonoTonicUpStack(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name      string
		args      args
		wantStack []int
	}{
		{
			name: "happy path",
			args: args{
				src: []int{3, 4, 2, 6, 4, 5, 2, 3},
			},
			wantStack: []int{6, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStack := GetMonoTonicUpStack(tt.args.src); !reflect.DeepEqual(gotStack, tt.wantStack) {
				t.Errorf("GetMonoTonicUpStack() = %v, want %v", gotStack, tt.wantStack)
			}
		})
	}
}
