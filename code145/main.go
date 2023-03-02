package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
	"reflect"
)

func postorderTraversal(root *treenode.TreeNode) []int {
	return bfs(root)
}

func bfs(node *treenode.TreeNode) []int {
	if node == nil {
		return nil
	}

	var ret []int
	var nodeStack []*treenode.TreeNode
	// 用于记录前一个加入到ret的节点的指针
	var preInsertNode *treenode.TreeNode
	for node != nil || len(nodeStack) > 0 {
		for node != nil {
			nodeStack = append(nodeStack, node)
			node = node.Left
		}

		node = nodeStack[len(nodeStack)-1]
		if node.Right != nil && node.Right != preInsertNode {
			node = node.Right
		} else {
			// 出栈
			nodeStack = nodeStack[:len(nodeStack)-1]
			// 将当前节点的值加入到返回结果中；并且把这个节点的指针用 preInsertNode 记录
			ret = append(ret, node.Val)
			preInsertNode = node
			node = nil
		}
	}
	return ret
}

func dfs(node *treenode.TreeNode, ret *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, ret)
	dfs(node.Right, ret)
	*ret = append(*ret, node.Val)
}

func main() {
	srcTree := &treenode.TreeNode{
		Val: 1,
		Left: &treenode.TreeNode{
			Val: 2,
			Left: &treenode.TreeNode{
				Val: 3,
			},
			Right: &treenode.TreeNode{
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
