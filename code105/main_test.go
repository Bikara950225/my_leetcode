package main

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "path1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: []int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTree := buildTree(tt.args.preorder, tt.args.inorder)
			got := gotTree.GetPreorder()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree2() = %v, want %v", got, tt.want)
			}
		})
	}
}
