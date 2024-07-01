package main

import . "my_leetcode/internal/tree_node"

func buildTree(preorder, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, len(inorder))
	for i, item := range inorder {
		inorderMap[item] = i
	}

	var dfsFunc func(preBegin, preEnd int, inBegin, inEnd int) *TreeNode
	dfsFunc = func(preBegin, preEnd int, inBegin, inEnd int) *TreeNode {
		if preBegin > preEnd || inBegin > inEnd {
			return nil
		}

		currVal := preorder[preBegin]
		currInorderIdx := inorderMap[currVal]

		// nextPreEnd - preBegin = currInorderIdx - inBegin
		nextPreEnd := currInorderIdx - inBegin + preBegin

		return &TreeNode{
			Val:   currVal,
			Left:  dfsFunc(preBegin+1, nextPreEnd, inBegin, currInorderIdx-1),
			Right: dfsFunc(nextPreEnd+1, preEnd, currInorderIdx+1, inEnd),
		}
	}
	tail := len(preorder) - 1
	return dfsFunc(0, tail, 0, tail)
}
