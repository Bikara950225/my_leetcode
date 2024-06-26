package main

import (
	"fmt"
	. "my_leetcode/internal/tree_node"
	"reflect"
)

func inorderTraversal(root *TreeNode) []int {
	return bfs(root)
}

func bfs(node *TreeNode) (ret []int) {
	if node == nil {
		return
	}

	var stack []*TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// pop
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		node = node.Right
	}
	return
}

func dfs(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, ret)
	*ret = append(*ret, node.Val)
	dfs(node.Right, ret)
}

func main() {
	srcTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	expectRet := []int{4, 2, 1, 3}

	ret1 := bfs(srcTree)
	if !reflect.DeepEqual(ret1, expectRet) {
		panic(fmt.Errorf("code94 error, not expect result: %+v",
			ret1,
		))
	}

	var ret2 []int
	dfs(srcTree, &ret2)
	if !reflect.DeepEqual(ret2, expectRet) {
		panic(fmt.Errorf("code94 error, not expect result: %+v",
			ret2,
		))
	}
}
