package list_node

import (
	"reflect"
	"testing"
)

func TestQuickCreateListNode(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				src: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := QuickCreateListNode(tt.args.src)
			gotList := got.GetList()
			if !reflect.DeepEqual(tt.want, gotList) {
				t.Errorf("TestQuickCreateListNode() error, not expect list: %v", gotList)
				return
			}
		})
	}
}
