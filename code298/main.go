package main

import (
	"fmt"
	"my_leetcode/internal/tools"
	treenode "my_leetcode/internal/tree_node"
)

func longestConsecutive(root *treenode.TreeNode) (ret int) {
	var dfsFunc func(n *treenode.TreeNode) int

	dfsFunc = func(n *treenode.TreeNode) int {
		if n == nil {
			return 0
		}

		leftL := dfsFunc(n.Left)
		if n.Left != nil && n.Val+1 != n.Left.Val {
			leftL = 0
		}

		rightL := dfsFunc(n.Right)
		if n.Right != nil && n.Val+1 != n.Right.Val {
			rightL = 0
		}

		maxL := tools.IntMax(leftL, rightL) + 1
		if maxL > ret {
			ret = maxL
		}
		return maxL
	}
	dfsFunc(root)

	return
}

func main() {
	tree1 := treenode.BuildSearchTreeWithPreorder([]int{1, 2, 5, 4})
	ret1 := longestConsecutive(tree1)
	if ret1 != 2 {
		panic(fmt.Sprintf("code298 error, ret1 != 2: %d", ret1))
	}
}
