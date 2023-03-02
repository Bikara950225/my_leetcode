package main

import (
	"my_leetcode/internal/tools"
	treenode "my_leetcode/internal/tree_node"
)

func minDepth(node *treenode.TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil {
		return minDepth(node.Right) + 1
	}
	if node.Right == nil {
		return minDepth(node.Left) + 1
	}
	return tools.IntMin(minDepth(node.Left), minDepth(node.Right))
}

func main() {

}
