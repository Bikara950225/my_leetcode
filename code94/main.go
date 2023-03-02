package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

func inorderTraversal(root *treenode.TreeNode) []int {
	return bfs(root)
}

func bfs(node *treenode.TreeNode) []int {
	if node == nil {
		return nil
	}

	var ret []int
	var nodeCache []*treenode.TreeNode
	for node != nil || len(nodeCache) > 0 {
		for node != nil {
			nodeCache = append(nodeCache, node)
			node = node.Left
		}
		// 出栈
		node = nodeCache[len(nodeCache)-1]
		nodeCache = nodeCache[:len(nodeCache)-1]

		ret = append(ret, node.Val)
		node = node.Right
	}
	return ret
}

func dfs(node *treenode.TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, ret)
	*ret = append(*ret, node.Val)
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
	expectRet := []int{4, 2 , 1, 3}

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
