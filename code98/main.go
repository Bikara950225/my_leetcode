package main

import (
	"fmt"
	"math"
	treenode "my_leetcode/internal/tree_node"
)

func isValidBST(root *treenode.TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(node *treenode.TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val >= upper || node.Val <= lower {
		return false
	}
	return dfs(node.Left, lower, node.Val) && dfs(node.Right, node.Val, upper)
}

func main() {
	src1 := &treenode.TreeNode{
		Val:   2,
		Left:  &treenode.TreeNode{Val: 1},
		Right: &treenode.TreeNode{Val: 3},
	}
	src1Ret := isValidBST(src1)
	if !src1Ret {
		panic(fmt.Errorf("code98 error, not expect result: %v", src1Ret))
	}

	src2 := &treenode.TreeNode{
		Val:   2,
		Left:  &treenode.TreeNode{Val: 2},
		Right: &treenode.TreeNode{Val: 2},
	}
	src2Ret := isValidBST(src2)
	if src2Ret {
		panic(fmt.Errorf("code98 error, not expect result: %v", src2Ret))
	}
}
