package main

import (
	"fmt"
	. "my_leetcode/internal/tree_node"
	"reflect"
)

func preorderTraversal(root *TreeNode) []int {
	return bfs(root)
}

// 广度优先遍历
func bfs(node *TreeNode) (ret []int) {
	if node == nil {
		return
	}

	var stack []*TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			ret = append(ret, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		// pop
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return
}

// 深度优先遍历
func dfs(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	*ret = append(*ret, node.Val)
	dfs(node.Left, ret)
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

	ret1 := bfs(srcTree)
	expectRet1 := []int{1, 2, 4, 3}
	if !reflect.DeepEqual(ret1, expectRet1) {
		panic(fmt.Errorf("code144 error, bfs() is not expect result: %+v", ret1))
	}

	var ret2 []int
	dfs(srcTree, &ret2)
	expectRet2 := []int{1, 2, 4, 3}
	if !reflect.DeepEqual(ret2, expectRet2) {
		panic(fmt.Errorf("code144 error, bfs() is not expect result: %+v", ret2))
	}
}
