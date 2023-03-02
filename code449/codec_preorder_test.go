package main

import (
	treenode "my_leetcode/internal/tree_node"
	"reflect"
	"testing"
)

var (
	testTree = &treenode.TreeNode{
		Val: 3,
		Left: &treenode.TreeNode{Val: 2},
		Right: &treenode.TreeNode{
			Val: 5,
			Left: &treenode.TreeNode{Val: 4},
		},
	}
)

func TestCodecPreorder_serialize(t *testing.T) {
	type args struct {
		root *treenode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "happy path",
			args: args{
				root: testTree,
			},
			want: `[3,2,5,4]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CodecPreorder{}
			if got := s.serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodecPreorder_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *treenode.TreeNode
	}{
		{
			name: "happy path",
			args: args{
				data: "[3,2,5,4]",
			},
			want: testTree,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CodecPreorder{}
			got := s.deserialize(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}