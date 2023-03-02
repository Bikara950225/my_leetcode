package tree_node

import (
	"reflect"
	"testing"
)

func TestTreeNode_GetPreorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "happy path1",
			fields: fields{
				Val: 10,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 11,
				},
			},
			want: []int{10, 5, 4, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := s.GetPreorder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_GetInorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "happy path1",
			fields: fields{
				Val: 10,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 11,
				},
			},
			want: []int{4, 5, 10, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := s.GetInorder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNode_GetPostorder(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "happy path1",
			fields: fields{
				Val: 10,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 11,
				},
			},
			want: []int{4, 5, 11, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := s.GetPostorder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildSearchTreeWithPreorder(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "happy path1",
			args: args{
				src: []int{
					2, 1, 3,
				},
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "happy path2",
			args: args{
				src: []int{
					10, 7, 4, 3, 5, 8, 11,
				},
			},
			want: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
				Right: &TreeNode{
					Val: 11,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildSearchTreeWithPreorder(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildSearchTreeWithPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
