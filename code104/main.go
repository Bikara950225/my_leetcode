package main

import (
	"my_leetcode/internal/tools"
	treenode "my_leetcode/internal/tree_node"
)

func maxDepth(root *treenode.TreeNode) int {
	if root == nil {
		return 0
	}
	return tools.IntMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {

}
