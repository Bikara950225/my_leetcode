package main

import (
	"fmt"
	. "my_leetcode/internal/tree_node"
	"reflect"
)

func rightSideView(root *TreeNode) (ret []int) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		ql := len(q)
		for i, curr := range q {
			if curr != nil {
				if i+1 == ql {
					ret = append(ret, curr.Val)
				}
				if curr.Left != nil {
					q = append(q, curr.Left)
				}
				if curr.Right != nil {
					q = append(q, curr.Right)
				}
			}
		}
		q = q[ql:]
	}
	return
}

func main() {
	srcTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	ret := rightSideView(srcTree)
	expectRet := []int{1, 3, 4}
	if !reflect.DeepEqual(ret, expectRet) {
		panic(fmt.Errorf("code199 error, not expect result: %+v", ret))
	}
}
