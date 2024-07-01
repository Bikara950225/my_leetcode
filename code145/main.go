package main

import (
	"fmt"
	. "my_leetcode/internal/tree_node"
	"reflect"
)

func postorderTraversal(root *TreeNode) []int {
	return bfs(root)
}

func bfs(node *TreeNode) (ret []int) {
	if node == nil {
		return
	}

	var stack []*TreeNode
	var mark *TreeNode
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		popNode := stack[len(stack)-1]
		if popNode.Right != nil && popNode.Right != mark {
			node = popNode.Right
		} else {
			stack = stack[:len(stack)-1]
			ret = append(ret, popNode.Val)
			mark = popNode
		}
	}
	return
}

func dfs(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, ret)
	dfs(node.Right, ret)
	*ret = append(*ret, node.Val)
}

func bfs2(node *TreeNode) (ret []int) {
	var stack []*TreeNode
	var mark *TreeNode
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		popNode := stack[len(stack)-1]
		if popNode.Right != nil && popNode.Right != mark {
			node = popNode.Right
		} else {
			stack = stack[:len(stack)-1]
			ret = append(ret, popNode.Val)
			mark = popNode
		}
	}
	return
}

func main() {
	srcTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	expectRet := []int{3, 4, 2, 1}

	ret1 := bfs(srcTree)
	if !reflect.DeepEqual(ret1, expectRet) {
		panic(fmt.Errorf("code145 error, not expect result: %+v", ret1))
	}

	var ret2 []int
	dfs(srcTree, &ret2)
	if !reflect.DeepEqual(ret2, expectRet) {
		panic(fmt.Errorf("code145 error, not expect result: %+v", ret2))
	}
}
