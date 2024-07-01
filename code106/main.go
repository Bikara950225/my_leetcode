package main

import (
	. "my_leetcode/internal/tree_node"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderIdxMap := make(map[int]int, len(inorder))
	for i, item := range inorder {
		inorderIdxMap[item] = i
	}

	var dfsFunc func(inBegin, inEnd int, postBegin, postEnd int) *TreeNode
	dfsFunc = func(inBegin, inEnd int, postBegin, postEnd int) *TreeNode {
		if inBegin > inEnd || postBegin > postEnd {
			return nil
		}

		val := postorder[postEnd]
		inorderIdx := inorderIdxMap[val]

		// nextPostEnd - postBegin = inorderIdx - 1 - inBegin
		nextPostEnd := inorderIdx - inBegin + postBegin - 1

		return &TreeNode{
			Val:   val,
			Left:  dfsFunc(inBegin, inorderIdx-1, postBegin, nextPostEnd),
			Right: dfsFunc(inorderIdx+1, inEnd, nextPostEnd+1, postEnd-1),
		}
	}
	tail := len(inorder) - 1
	return dfsFunc(0, tail, 0, tail)
}
