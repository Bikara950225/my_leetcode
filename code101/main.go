package main

import (
	"fmt"
	treenode "my_leetcode/internal/tree_node"
)

func isSymmetric(root *treenode.TreeNode) bool {
	if root == nil {
		return false
	}

	queue := []*treenode.TreeNode{
		root.Left, root.Right,
	}
	for len(queue) > 0 {
		if !check(queue) {
			return false
		}
		currLen := len(queue)
		for i := 0; i < currLen; i++ {
			if queue[i] == nil {
				continue
			}
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		queue = queue[currLen:]
	}
	return true
}

func check(queue []*treenode.TreeNode) bool {
	i, j := 0, len(queue)-1
	for i < j {
		iNode := queue[i]
		jNode := queue[j]
		if iNode == nil && jNode != nil {
			return false
		}
		if iNode != nil && jNode == nil {
			return false
		}
		if iNode != nil && jNode != nil && iNode.Val != jNode.Val {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	srcTree1 := &treenode.TreeNode{
		Val: 1,
		Left: &treenode.TreeNode{
			Val:   2,
			Right: &treenode.TreeNode{Val: 3},
		},
		Right: &treenode.TreeNode{
			Val:   2,
			Right: &treenode.TreeNode{Val: 3},
		},
	}
	if isSymmetric(srcTree1) {
		panic(fmt.Errorf("code101 error, true"))
	}

	srcTree2 := &treenode.TreeNode{
		Val: 1,
		Left: &treenode.TreeNode{
			Val:   2,
			Right: &treenode.TreeNode{Val: 3},
		},
		Right: &treenode.TreeNode{
			Val:  2,
			Left: &treenode.TreeNode{Val: 3},
		},
	}
	if !isSymmetric(srcTree2) {
		panic(fmt.Errorf("code101 error, false"))
	}
}
