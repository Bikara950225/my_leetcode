package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
)

var ret bool = true

func isBalanced(root *treenode.TreeNode) bool {
	defer func() {
		ret = true
	}()

	getNodeHight(root)
	return ret
}

func getNodeHight(node *treenode.TreeNode) int {
	if node == nil {
		return 0
	}

	leftH := getNodeHight(node.Left)
	rightH := getNodeHight(node.Right)

	if leftH > rightH && leftH-rightH > 1 {
		ret = false
	}
	if leftH < rightH && rightH-leftH > 1 {
		ret = false
	}
	return max(leftH, rightH) + 1
}

func max(src1, src2 int) int {
	if src1 > src2 {
		return src1
	}
	return src2
}

func main() {
	src1 := &treenode.TreeNode{
		Val:  3,
		Left: &treenode.TreeNode{Val: 9},
		Right: &treenode.TreeNode{
			Val: 20,
			Left: &treenode.TreeNode{
				Val: 15,
			},
			Right: &treenode.TreeNode{
				Val: 7,
			},
		},
	}
	ret1 := isBalanced(src1)
	if !ret1 {
		panic(fmt.Errorf("code110 error, not expect: %t", ret1))
	}

	src2 := treenode.BuildSearchTreeWithPreorder([]int{10, 7, 4, 3, 5, 8, 11})
	ret2 := isBalanced(src2)
	if ret2 {
		panic(fmt.Errorf("code110 error, not expect: %t", ret2))
	}
}
