package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
)

func sumNumbers(root *treenode.TreeNode) (ret int) {
	var dfsFunc func(node *treenode.TreeNode, subRet int)
	dfsFunc = func(node *treenode.TreeNode, subRet int) {
		if node == nil {
			return
		}
		subRet = (10 * subRet) + node.Val
		if node.Left == nil && node.Right == nil {
			ret += subRet
			return
		}
		dfsFunc(node.Left, subRet)
		dfsFunc(node.Right, subRet)
	}
	dfsFunc(root, 0)

	return
}

func main() {
	ret1 := sumNumbers(&treenode.TreeNode{
		Val:   1,
		Left:  &treenode.TreeNode{Val: 2},
		Right: &treenode.TreeNode{Val: 3},
	})
	if ret1 != 25 {
		panic(fmt.Errorf("code129 error, ret1 != 25: %d", ret1))
	}
}
