package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

func preorderTraversal(root *treenode.TreeNode) []int {
	return bfs(root)
}

// 广度优先遍历
func bfs(node *treenode.TreeNode) []int {
	if node == nil {
		return nil
	}

	var ret []int
	var nodeStack []*treenode.TreeNode
	for node != nil || len(nodeStack) > 0 {
		for node != nil {
			ret = append(ret, node.Val)
			nodeStack = append(nodeStack, node)
			node = node.Left
		}
		// 出栈
		node = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		node = node.Right
	}
	return ret
}

// 深度优先遍历
func dfs(node *treenode.TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	*ret = append(*ret, node.Val)
	dfs(node.Left, ret)
	dfs(node.Right, ret)
}

func main() {
	srcTree := &treenode.TreeNode{
		Val: 1,
		Left: &treenode.TreeNode{
			Val: 2,
			Left: &treenode.TreeNode{
				Val: 4,
			},
		},
		Right: &treenode.TreeNode{
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
