package main

import (
	"my_leetcode/internal/tools"
	. "my_leetcode/internal/tree_node"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return tools.IntMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {

}
